package main

import "fmt"

func main() {

	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}

	println("Person 1:", p1.Name, "Age:", p1.Age)
	println("Person 2:", p2.Name, "Age:", p2.Age)

	// anonymous struct
	var p3 struct {
		Name string
		Age  int
	}

	p3.Name = "Charlie"
	fmt.Println("Person 3:", p3.Name, "Age:", p3.Age)

	fmt.Println(p1 == p3)
}
