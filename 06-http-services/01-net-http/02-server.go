package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

var products []Product = []Product{
	{101, "Pen", 10},
	{102, "Pencil", 5},
	{103, "Marker", 50},
}

type AppServer struct {
	routes map[string]http.HandlerFunc
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (appServer *AppServer) AddRoute(pattern string, handler http.HandlerFunc) {
	appServer.routes[pattern] = handler
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := appServer.routes[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

// application specific logic

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "data serialization error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newProduct); err != nil {
			http.Error(w, "data serialization error", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "method not implemented", http.StatusMethodNotAllowed)
	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "{\"response\" : \"All the customers will be served\"}")
}

// middlewares
func logMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("trace-id : %v - %s - %s\n", r.Context().Value("trace-id"), r.Method, r.URL.Path)
		next(w, r)
	}
}

func jsonMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	}
}

func traceMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		traceId := rand.Intn(100)
		traceCtx := context.WithValue(r.Context(), "trace-id", traceId)
		reqWithTraceId := r.WithContext(traceCtx)
		next(w, reqWithTraceId)
	}
}

func timeoutMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeoutCtx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		reqWithTimeoutCtx := r.WithContext(timeoutCtx)
		go func() {
			<-reqWithTimeoutCtx.Context().Done()
			if reqWithTimeoutCtx.Context().Err() == context.DeadlineExceeded {
				fmt.Println("[timeoutMiddleware] sending timeout response")
				http.Error(w, `{"response" : "request timed out"}`, http.StatusRequestTimeout)
			}
		}()
		next(w, reqWithTimeoutCtx)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	prepareIndexResponse(15 * time.Second)
	fmt.Println("prepared response")
	select {
	case <-r.Context().Done():
		return
	default:
		var response = make(map[string]any)
		response["trace-id"] = r.Context().Value("trace-id")
		response["response"] = "Hello,World!"
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "data serialization error", http.StatusInternalServerError)
		}
	}
}

func prepareIndexResponse(d time.Duration) {
	time.Sleep(d)
}

func main() {
	appServer := NewAppServer()
	appServer.AddRoute("/", timeoutMiddleware(traceMiddleware(logMiddleware(jsonMiddleware(IndexHandler)))))
	appServer.AddRoute("/products", timeoutMiddleware(traceMiddleware(logMiddleware(jsonMiddleware(ProductsHandler)))))
	appServer.AddRoute("/customers", timeoutMiddleware(traceMiddleware(logMiddleware(jsonMiddleware(CustomersHandler)))))
	http.ListenAndServe(":8080", appServer)
}
