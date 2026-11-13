package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	y := 10
	fmt.Println(y)
	// fmt := "oops"
	// fmt.Println(fmt)
}
