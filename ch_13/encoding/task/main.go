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
}

type ProductList struct {
	Products []Product `json:"products"`
}

func main() {

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("xato:", err)
		return
	}

	var list ProductList
	err = json.Unmarshal(data, &list)
	if err != nil {
		fmt.Println("xato:", err)
		return
	}
	fmt.Println("O'qildi:", len(list.Products), "ta product")

	// create file
	out, err := os.Create("products.jsonl")
	if err != nil {
		fmt.Println("xato:", err)
		return
	}

	enc := json.NewEncoder(out)
	for _, product := range list.Products {
		if err := enc.Encode(product); err != nil {
			fmt.Println("yozishda xato:", err)
			return
		}
	}

	out.Close()

	in, err := os.Open("products.jsonl")
	if err != nil {
		fmt.Println("xato:", err)
		return
	}
	defer in.Close()

	dec := json.NewDecoder(in)
	count := 0
	for dec.More() { // yana JSON bormi?
		var p Product // har safar yangi, toza struct
		if err := dec.Decode(&p); err != nil {
			fmt.Println("o'qishda xato:", err)
			return
		}
		count++
		fmt.Printf("%d. %s — %.2f\n", p.ID, p.Name, p.Price)
	}
	fmt.Println("Stream'dan o'qildi:", count, "ta")

	raw, _ := os.ReadFile("test.jsonl")
	var one Product
	err = json.Unmarshal(raw, &one)
	if err != nil {
		fmt.Println("errir", err)
		return
	}
	fmt.Println("Unmarshal natijasi:", string(raw))
}
