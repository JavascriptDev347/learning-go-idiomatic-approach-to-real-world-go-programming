package main

import (
	"fmt"
	"sync"
)

func main1() {
	runThreeGoroutines()
}

func runThreeGoroutines() {

	// chanel ma'lumotlar almashish uchun
	ch := make(chan int)

	// 2 ta yozish bor ularni kuzatish uchun wg
	var wg sync.WaitGroup
	wg.Add(2)

	// 1- goroutine 10 ta sonni yozadi
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 11; i <= 20; i++ {
			ch <- i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	// 3- goroutine tugashi uchun done
	done := make(chan struct{})

	go func() {
		defer close(done)
		for v := range ch {
			println(v)
		}
	}()

	<-done
	fmt.Println("Barcha sonlar chop etildi, dastur xavfsiz yakunlandi!")
}
