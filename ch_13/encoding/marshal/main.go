package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Discount     int64  `json:"discount,omitempty"`
	InternalNote string `json:"-"`
}

func main() {

	p := Product{
		ID:           "123",
		Name:         "Widget",
		Price:        1000,
		Discount:     0,
		InternalNote: "This is an internal note",
	}

	// data, err := json.Marshal(p)
	// if err != nil {
	// 	fmt.Println("Marshallingda xatolik berdi")
	// 	return
	// }

	// println(string(data))

	// chiroyli formatlash uchun MarshalIndent ishlatamiz
	data, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(data))

}
