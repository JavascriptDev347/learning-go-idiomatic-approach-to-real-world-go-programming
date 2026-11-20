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
		fmt.Println("Fayl yaratishda xato:", err)
		return
	}
	defer dst.Close()

	buf := make([]byte, 1024)
	totalReads := 0
	totalWrites := 0

	for {
		n, readErr := src.Read(buf)
		totalReads += n

		m, writeErr := dst.Write(buf[:n])
		totalWrites += m

		if writeErr != nil {
			fmt.Println("Yozishda xato:", writeErr)
			break
		}
		if readErr == io.EOF {
			fmt.Println("O'qish tugadi")
			break
		}
		if readErr != nil {
			fmt.Println("O'qishda xato:", readErr)
			break
		}
	}

	fmt.Println("Total reads:", totalReads)
	fmt.Println("Total writes:", totalWrites)

}
