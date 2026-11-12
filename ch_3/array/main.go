package main

import "fmt"

func main() {
	var x [3]int // the array has 3 elements that all of them 0. because it inizialized default value of int
	x[2] = 90
	fmt.Println(x)
	// i := 6
	// fmt.Println(x[i])// error

	var y = [...]int{0, 0, 90}

	fmt.Println(x == y) // true

}
