package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	src, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Faylni ochishda xato:", err)
		return
	}
	defer src.Close()

	dst, err := os.Create("copy.txt")
	if err != nil {
		fmt.Println("Faylni yaratishda xato:", err)
		return
	}
	defer dst.Close()

	// copy
	written, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Faylni nusxalashda xato:", err)
		return
	}

	fmt.Printf("%d bayt nusxalandi.\n", written)
}
