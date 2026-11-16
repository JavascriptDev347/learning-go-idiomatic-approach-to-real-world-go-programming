package main

import "fmt"

func main() {

	r, err := checkEven(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(r)

}

func checkEven(a int) (int, error) {
	if a%2 != 0 {
		return 0, fmt.Errorf("%d is not an even number", a)
	}
	return a * 2, nil
}
