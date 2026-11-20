package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount,omitempty"`
}

func main() {
	input := []byte(`{"id":"p2","name":"Noutbuk","price":9000000,"discount":500000}`)
	var p Product
	err := json.Unmarshal(input, &p)
	if err != nil {
		fmt.Println("xato:", err)
		return
	}
	fmt.Println(p.Name, p.Price, p.Discount)
}
