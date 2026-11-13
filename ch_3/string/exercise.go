package main

import "fmt"

func main() {
	var s string = "Hello, World! 😎"
	var s2 = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s[6])

	var bs = []byte(s)
	var rs = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
