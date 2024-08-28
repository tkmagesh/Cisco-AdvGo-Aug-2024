/*
interfaces
  - contracts
  - "implicitly" implemented
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// ver 1.0
type Product struct {
	Id   int
	Name string
	Cost float64
}

type ProductService struct {
	products []Product
}

func NewProductService() *ProductService {
	return &ProductService{
		products: []Product{
			{101, "Pen", 10},
			{102, "Pencil", 5},
			{103, "Marker", 50},
		},
	}
}

func (ps *ProductService) GetAll() []Product {
	return ps.products
}

// ver 2.0
type ProductCSVSerializer struct {
	/* ? */
}

func NewProductCSVSerializer(/* ? */) *ProductCSVSerializer {
	return &ProductCSVSerializer{
		/* ? */
	}
}

func (pSerializer *ProductCSVSerializer) Serialize() string {
	builder := strings.Builder{}
	var products []Product
	products := // get the products from the source
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%d,%q,%0.2f\n", p.Id, p.Name, p.Cost))
	}
	return builder.String()
}

func main() {
	productSerializer := // create the instance of ProductCSVSerializer
	fmt.Println(productSerializer.Serialize())
}
