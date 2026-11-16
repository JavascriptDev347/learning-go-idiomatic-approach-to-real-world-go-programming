package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)

		go worker(i, &wg)
	}
	wg.Wait()

}
