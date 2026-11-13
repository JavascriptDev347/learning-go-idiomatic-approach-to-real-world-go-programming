package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// fmt.Println(s1==s2) // invalid operation
	fmt.Println(slices.Equal(s1, s2))
	fmt.Println(CustomEqual(s1, s2))

	// var x []int
	// fmt.Println(x)
	// x = append(x, 10)
	// fmt.Println(x)

	// var y = make([]int, 255, 256)
	// fmt.Println("Default", y, len(y), cap(y))
	// y = append(y, 10)
	// fmt.Println("Append 10", y, len(y), cap(y))
	// y = append(y, 90)
	// fmt.Println("Append 90", y, len(y), cap(y))
	// z := make([]int, 10, 80) // 10 ta element initialize qilindi value = 0
	// fmt.Println(z)
	// i := []int{1, 2, 3, 4, 54, 90}
	// fmt.Println(i, len(i))
	// clear(i)
	// fmt.Println(i, len(i))

	// slicing slice
	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[1:]
	// d := x[1:3]
	// e := x[:]
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)
	// fmt.Println("d:", d)
	// fmt.Println("e:", e)

	// fmt.Println("x[1:3]:", x[1:3])
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
}

func CustomEqual(s1, s2 []int) bool {
	// 1-qadam
	if len(s1) != len(s2) {
		return false
	}

	// 2-qadam
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
