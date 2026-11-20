package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	// Discount float64 `json:"discount,omitempty"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

func main() {
	// input := []byte(`{"id":"p2","name":"Noutbuk","price":9000000,"discount":500000}`)
	// var p Product
	// err := json.Unmarshal(input, &p)
	// if err != nil {
	// 	fmt.Println("xato:", err)
	// 	return
	// }
	// fmt.Println(p.Name, p.Price, p.Discount)

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("xato:", err)
		return
	}
	var p ProductList
	err = json.Unmarshal(data, &p)

	if err != nil {
		fmt.Println("xato:", err)
		return
	}

	var total float64
	for _, product := range p.Products {
		total += product.Price
		fmt.Println(product.Name, product.Price)
	}

	fmt.Printf("Total price: %.2f\n", total)
}
