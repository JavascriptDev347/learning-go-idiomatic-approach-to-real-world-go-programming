package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	p1 := MakePerson("John", "Doe", 30)
	p2 := MakePersonPointer("Jane", "Smith", 25)

	println(p1.FirstName, p1.LastName, p1.Age) // prints John Doe 30
	println(p2.FirstName, p2.LastName, p2.Age) // prints Jane Smith 25

	p := Person{FirstName: "Jon", LastName: "Bodner", Age: 40}
	total := 10_000_000
	// 1-holat: capacity ko'rsatmasdan append qilish (slays o'sib boradi)
	start1 := time.Now()
	var people1 []Person
	for i := 0; i < total; i++ {
		people1 = append(people1, p)
	}
	duration1 := time.Since(start1)
	fmt.Printf("Dynamic slice (append): %v/n", duration1)

	// 2-holat: Oldindan sig'im (capacity) ajratilgan slays
	start2 := time.Now()
	people2 := make([]Person, 0, total)
	for i := 0; i < total; i++ {
		people2 = append(people2, p)
	}
	duration2 := time.Since(start2)
	fmt.Printf("Pre-allocated slice: %v\\n", duration2)

}
