package main

import "time"

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	// Items       []Item    `json:"items"`
}

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Discount     int64  `json:"discount,omitempty"`
	InternalNote string `json:"-"`
}

func main() {

}
