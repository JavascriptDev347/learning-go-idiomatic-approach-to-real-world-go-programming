package main

import "fmt"

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("First:", val)
	}(a)

	a = 20
	defer func(val int) {
		fmt.Println("Second:", val)
	}(a)

	a = 30
	fmt.Println("Exitting a:", a)
	return a
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func multiplier(k int) func(int) int {
	return func(x int) int {
		return k * x
	}
}

func main() {

	result := deferExample()
	fmt.Println("Result:", result)
	// c1 := counter()
	// println(c1(), c1(), c1())

	// c2 := counter()
	// println(c2())
	// println(c1())

	// double := multiplier(2)
	// triple := multiplier(3)
	// fmt.Println(double(5), triple(5))

	// fourBase := makeMult(4)
	// sixBase := makeMult(6)

	// for i := 0; i < 10; i++ {
	// 	println(fourBase(i), " ", sixBase(i))
	// }
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
