package main

import "fmt"

func main() {
	// s := make([]int, 3, 4)
	// fmt.Println("Birinchi xolat:", len(s), cap(s))
	// s = append(s, 6)
	// fmt.Println("Ikkinchi xolat:", len(s), cap(s))

	// s = append(s, 7, 8, 9, 0, 12)
	// fmt.Println("Uchinchi xolat:", len(s), cap(s))

	s2 := make([]int, 3, 4)
	fmt.Println("Birinchi xolat:", len(s2), cap(s2))
	s2 = append(s2, 6, 6, 6, 6, 6, 5, 54, 4)
	fmt.Println("Ikkinchi xolat:", len(s2), cap(s2))

}
