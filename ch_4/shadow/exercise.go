package main

import (
	"fmt"
)

func main() {
	x := 10
	if x > 5 {
		fmt.Println("Birinchi: ", x)
		x := 5
		fmt.Println("Ikkinchi: ", x)
	}
	fmt.Println("Uchinchi: ", x)

	// 	y := 10
	// 	fmt.Println(y)
	// 	// fmt := "oops"
	// 	// fmt.Println(fmt)

	//	a := rand.Intn(10)
	//	for a < 100 {
	//		if a%5 == 0 {
	//			goto done
	//		}
	//		a = a*2 + 1
	//	}
	//	fmt.Println("do something when the loop completes normally")
	//
	// done:
	//
	//	fmt.Println("do complicated stuff no matter why we left the loop")
	//	fmt.Println(a)
}
