package main

import "time"

type Person struct {
	Name      string
	Email     string
	DateAdded time.Time
}

func NewPerson(username, email string) *Person {
	return &Person{
		Name:      username,
		Email:     email,
		DateAdded: time.Now(),
	}
}

func main() {}
