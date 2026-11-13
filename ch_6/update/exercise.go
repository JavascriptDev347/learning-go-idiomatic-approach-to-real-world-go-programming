package main

import "fmt"

func main() {
	// var f *int
	p := 10
	fmt.Println(p) // prints <nil>
	failedUpdate(&p)
	fmt.Println(p) // prints 45
	update(&p)
	fmt.Println(p) // prints <nil>

}

func failedUpdate(x *int) {
	x2 := 20
	x = &x2
}

func update(x *int) {
	*x = 45
}
