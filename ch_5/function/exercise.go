package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name  string
	Email string
	Age   int
}

func main() {

	// var v, err = divide(10, 0)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Result:", v)
	// }

	// var a = gretting("Alice")
	// fmt.Println(a)

	people := []Person{
		{Name: "Alice", Email: "alice@example.com", Age: 30},
		{Name: "Bob", Email: "bob@example.com", Age: 25},
		{Name: "Charlie", Email: "charlie@example.com", Age: 35},
	}
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)

}

func gretting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func multiply(a, b int) int {
	return a * b
}
