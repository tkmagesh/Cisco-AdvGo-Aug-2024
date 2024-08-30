package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintln(w, "{\"response\" : \"Hello, World!\"}")
}

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
		log.Printf("%s - %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func main() {
	appServer := NewAppServer()
	appServer.AddRoute("/", logMiddleware(IndexHandler))
	appServer.AddRoute("/products", logMiddleware(ProductsHandler))
	appServer.AddRoute("/customers", logMiddleware(CustomersHandler))
	http.ListenAndServe(":8080", appServer)
}
