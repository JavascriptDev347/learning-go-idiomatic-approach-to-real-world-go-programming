package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	BaseURL     = "https://fakestoreapi.com"
	ProductsURL = BaseURL + "/products"
	UsersURL    = BaseURL + "/users"
	CartsURL    = BaseURL + "/carts"
)

type result[T any] struct {
	data T
	err  error
}

func main() {

	start := time.Now()
	var (
		wg               sync.WaitGroup
		products         []Product
		users            []User
		carts            []Cart
		pErr, uErr, cErr error
	)

	wg.Add(3)
	go func() { defer wg.Done(); products, pErr = fetchProducts() }()
	go func() { defer wg.Done(); users, uErr = fetchUsers() }()
	go func() { defer wg.Done(); carts, cErr = fetchCarts() }()
	wg.Wait()
	fmt.Println("Products:", len(products), "Users:", len(users), "Carts:", len(carts))

	if pErr != nil {
		fmt.Println("Products error:", pErr)
	}
	if uErr != nil {
		fmt.Println("Users error:", uErr)
	}
	if cErr != nil {
		fmt.Println("Carts error:", cErr)
	}

	fmt.Println("Vaqt:", time.Since(start))

}

// productsCh := make(chan result[[]Product], 1)
// usersCh := make(chan result[[]User], 1)
// cartsCh := make(chan result[[]Cart], 1)

// go func() { p, err := fetchProducts(); productsCh <- result[[]Product]{p, err} }()
// go func() { u, err := fetchUsers(); usersCh <- result[[]User]{u, err} }()
// go func() { c, err := fetchCarts(); cartsCh <- result[[]Cart]{c, err} }()

// p, u, c := <-productsCh, <-usersCh, <-cartsCh

// for _, err := range []error{p.err, u.err, c.err} {
// 	if err != nil {
// 		fmt.Println("Xatolik:", err)
// 	}
// }

// fmt.Println("Products:", len(p.data))
// fmt.Println("Users:", len(u.data))
// fmt.Println("Carts:", len(c.data))

func fetchProducts() ([]Product, error) {
	var products []Product
	if err := fetch(ProductsURL, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func fetchUsers() ([]User, error) {
	var users []User
	if err := fetch(UsersURL, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func fetchCarts() ([]Cart, error) {
	var carts []Cart
	if err := fetch(CartsURL, &carts); err != nil {
		return nil, err
	}
	return carts, nil
}
