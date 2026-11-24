package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Xatolik fiel ochishda", err)
	}
	buf := make([]byte, 4)

	for {
		n, err := f.Read(buf)
		if n > 0 {
			fmt.Printf("O'qildi: %d bayt | Matn: %q\n", n, string(buf[:n]))
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("o'qish muvafaqiyatli EOF ")
				break
			}
			fmt.Println("Xatolik yuz berdi:", err)
			return
		}
	}
}
