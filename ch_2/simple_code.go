package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Boolean
	var flag bool
	flag = false
	fmt.Println(flag)

	// integer type
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint 64
	//
	// byte -> uint8
	var i int //-128 to 127
	fmt.Println(i)

	// float 32 and 64
	//
	// 	a := 0.1
	// 	b := 0.2
	//
	// 	fmt.Println(a + b)

	var a float64 = 0.1
	var b float64 = 0.2

	fmt.Printf("%.20f\n", a)
	fmt.Printf("%.20f\n", b)
	fmt.Printf("%.20f\n", a+b)

	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)

	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

}
