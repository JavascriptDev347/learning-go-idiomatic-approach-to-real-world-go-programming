package main

import "fmt"

const w int = 90

func main() {

	const y = "hello"

	fmt.Println(w)
	fmt.Println(y)

	//ch_2\simple_code.go:14:2: cannot assign to w (neither addressable nor a map index expression)
	// ch_2\simple_code.go:15:2: cannot assign to y (neither addressable nor a map index expression)
	// w = w + 1
	// y = "bye"

	fmt.Println(w)
	fmt.Println(y)

	// 10 and 30 are not read.
	x := 10
	x = 20
	fmt.Println(x)
	x = 30
	fmt.Println(x)

}

// Boolean
// var flag bool
// flag = false
// fmt.Println(flag)

// integer type
// int8, int16, int32, int64
// uint8, uint16, uint32, uint 64
//
// byte -> uint8
// var i int //-128 to 127
// fmt.Println(i)

// float 32 and 64
//
// 	a := 0.1
// 	b := 0.2
//
// 	fmt.Println(a + b)

// 	var a float64 = 0.1
// 	var b float64 = 0.2
//
// 	fmt.Printf("%.20f\n", a)
// 	fmt.Printf("%.20f\n", b)
// 	fmt.Printf("%.20f\n", a+b)

// x := complex(2.5, 3.1)
// y := complex(10.2, 2)
// 	fmt.Println(x + y)
// 	fmt.Println(x - y)
// 	fmt.Println(x * y)
// 	fmt.Println(x / y)
//
// 	fmt.Println(real(x))
// 	fmt.Println(imag(x))
// 	fmt.Println(cmplx.Abs(x))

// Strings
// var s string = "Salom"
// s[0] = "P" stringni ichki baytlarini o'zgartirib bo'lmaydi

// 	s := "hello world !!!"
//
// 	bt := []byte(s)
// 	bt[0] = 'H'
// 	s = string(bt)
// 	fmt.Println(s)
//
// 	var r = 'R'
// 	fmt.Println(r)
// 	fmt.Printf("%c\n", r)
// 	fmt.Printf("%T\n", r)
