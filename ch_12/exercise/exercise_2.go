package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 1; i <= 10; i++ {
			ch2 <- i * 10
		}
	}()

	closedChannels := 0

	for closedChannels < 2 {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				closedChannels++
				continue
			}
			fmt.Printf("[Goroutina 1] dan keldi: %d\\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				closedChannels++
				continue
			}
			fmt.Printf("[Goroutina 2] dan keldi: %d\\n", v)
		}
	}
	fmt.Println("Barcha 20 ta son o'qildi, dastur xavfsiz yakunlandi!")
}
