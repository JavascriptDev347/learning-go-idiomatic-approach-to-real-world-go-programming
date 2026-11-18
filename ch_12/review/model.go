package main

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

type Cart struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userId"`
	Date     string `json:"date"`
	Products []struct {
		ProductID int `json:"productId"`
		Quantity  int `json:"quantity"`
	} `json:"products"`
}

type User struct {
	Address struct {
		Geolocation Geolocation `json:"geolocation"`
		City        string      `json:"city"`
		Street      string      `json:"street"`
		Number      int         `json:"number"`
		Zipcode     string      `json:"zipcode"`
	} `json:"address"`
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"name"`
	Phone string `json:"phone"`
}

type Geolocation struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}
