package main

import "fmt"

type person struct {
	name string
	age  int
}

func modifyFails(n string, a int, p person) {
	a = a * 12
	n = "HEllo"
	p.name = "BOb"

}
func main() {

	p := person{}
	i := 2
	n := "Russel"
	modifyFails(n, i, p)
	fmt.Printf("n: %s, i: %d, p: %v\n", n, i, p) //

}
