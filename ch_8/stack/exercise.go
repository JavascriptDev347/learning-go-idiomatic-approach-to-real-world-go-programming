package main

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return val, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() {

	var intStack Stack[int]
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack) // Output: {vals:[1 2 3]}
	fmt.Println(intStack.Contains(1))
	val, ok := intStack.Pop()
	if ok {
		println(val) // Output: 3
	}

	fmt.Println(intStack) // Output: {vals:[1 2]}

	var stringStack Stack[string]
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack) // Output: {vals:[hello world]}

	strVal, strOk := stringStack.Pop()
	if strOk {
		println(strVal) // Output: world
	}
}
