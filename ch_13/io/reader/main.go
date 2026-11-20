package main

import (
	"fmt"
	"io"
	"os"
)

// type Reader interface {
// 		Read(p []byte) (n int, err error)
// 	}

// 	type Writer interface {
// 		Write(p []byte) (n int, err error)
// 	}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("Fayl hajmi:", info.Size())
	defer f.Close()

	buf := make([]byte, 1024)
	var counter int = 0
	var totalBytesRead int = 0
	bytes := make([]byte, 0, 8)
	for {
		n, err := f.Read(buf)
		counter = counter + 1
		totalBytesRead += n
		bytes = append(bytes, buf[:n]...)

		fmt.Println("quti hajmi:", len(buf))
		fmt.Println("nechta solindi (n):", n)
		fmt.Printf("butun quti: %q\n", buf)
		fmt.Printf("faqat yangilari buf[:n]: %q\n", buf[:n])
		fmt.Println("----------")

		if err == io.EOF {
			fmt.Println("Tugadi")
			break
		}
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Counter", counter)
	fmt.Println("total bytes va file hajmi:", totalBytesRead, info.Size())
	fmt.Println("Total bytes read:", totalBytesRead)
	fmt.Println(string(bytes))

}

// r, err := countLetters(strings.NewReader("Hello, World!"))
// if err != nil {
// 	panic(err)
// }
// for k, v := range r {
// 	println(k, v)
// }

// r := strings.NewReader("s")

// // Quti: 5 ta joy
// buf := make([]byte, 3)

// for {
// 	n, err := r.Read(buf) // do'konga bordik

// 	fmt.Println("quti hajmi:", len(buf))
// 	fmt.Println("nechta solindi (n):", n)
// 	fmt.Printf("butun quti: %q\n", buf[:n])
// 	fmt.Printf("faqat yangilari buf[:n]: %q\n", buf[:n])
// 	fmt.Println("----------")

// 	if err == io.EOF { // do'konchi: "tugadi"
// 		fmt.Println("tugadi")
// 		break
// 	}
// }

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		fmt.Println("Read", n, "bytes")
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
