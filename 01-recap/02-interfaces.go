/*
interfaces
  - contracts
  - "implicitly" implemented
*/
package main

import (
	"fmt"
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
/* Violation of Dependency Inversion Principle */
/*
type ProductCSVSerializer struct {
	productService *ProductService
}

func NewProductCSVSerializer(ps *ProductService) *ProductCSVSerializer {
	return &ProductCSVSerializer{
		productService: ps,
	}
}

func (pSerializer *ProductCSVSerializer) Serialize() string {
	builder := strings.Builder{}
	var products []Product
	products = pSerializer.productService.GetAll()
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%d,%q,%0.2f\n", p.Id, p.Name, p.Cost))
	}
	return builder.String()
}
*/

type IProductService interface {
	GetAll() []Product
}

type ProductCSVSerializer struct {
	productService IProductService
}

func NewProductCSVSerializer(ps IProductService) *ProductCSVSerializer {
	return &ProductCSVSerializer{
		productService: ps,
	}
}

func (pSerializer *ProductCSVSerializer) Serialize() string {
	builder := strings.Builder{}
	var products []Product
	products = pSerializer.productService.GetAll()
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%d,%q,%0.2f\n", p.Id, p.Name, p.Cost))
	}
	return builder.String()
}

func main() {
	ps := NewProductService()
	productSerializer := NewProductCSVSerializer(ps)
	fmt.Println(productSerializer.Serialize())
}
