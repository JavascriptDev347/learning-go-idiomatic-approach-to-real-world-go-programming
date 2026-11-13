package main

import (
	"fmt"
	"maps"
)

func main() {

	// reading and writing in map
	totalWins := map[string]int{
		"Rafael Nadal":   92,
		"Novak Djokovic": 87,
		"Roger Federer":  103,
	}

	totalWins["Andy Murray"] = 46
	fmt.Println(totalWins["Russel"]) // 0
	fmt.Println(totalWins["Rafael Nadal"])

	//comma ok
	v, ok := totalWins["Russel"]
	fmt.Println(v, ok) // 0 false
	v, ok = totalWins["Rafael Nadal"]
	fmt.Println(v, ok) // 92 true

	// deleting from map
	delete(totalWins, "Andy Murray")
	fmt.Println(totalWins["Andy Murray"]) // 0

	// iterating over map
	for k, v := range totalWins {
		fmt.Println(k, v)
	}

	m1 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	m2 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	// fmt.Println(m1 == m2) // invalid operation: m1 == m2 (map can only be compared to nil)
	fmt.Println(maps.Equal(m1, m2))

	// using map as a set
	intSet := map[int]bool{}
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(intSet)

}
