package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := []byte("Hello, World!")

	n, err := file.Write(data)
	m, _ := file.Write(data)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(n, m)
}
