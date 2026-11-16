package main

import "fmt"

func main() {
	fmt.Println("---First exercise---")
	fmt.Println(Double(29.2))
	fmt.Println(Double(90))

	fmt.Println("---Second exercise---")
	PrintVal(PrintableInt(42))
	PrintVal(PrintableFloat(3.14))
	
}

// exercise 1

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Double[T Number](val T) T {
	return val * 2
}

// Exercise 2
type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type PrintableInt int

func (p PrintableInt) String() string {
	return fmt.Sprintf("PrintableInt qiymati: %d", p)
}

type PrintableFloat float64

func (p PrintableFloat) String() string {
	return fmt.Sprintf("PrintableFloat qiymati: %.2f", p)
}

func PrintVal[T Printable](val T) {
	fmt.Println(val)
}
