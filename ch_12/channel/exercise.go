package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d: Job %d started\n", id, job)

		time.Sleep(time.Second)

		result := job * 2

		fmt.Printf("Worker %d: Job %d completed\n", id, job)

		results <- result
	}
}

func main() {

	// 1. Buffered channels yaratamiz
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup

	// 2. Uchta worker goroutine ishga tushiramiz
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// 3. Jobs channel'iga 5 ta vazifa yuboramiz
	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	// 4. Boshqa vazifalar yuborilmasligini bildiramiz
	close(jobs)

	// Barcha worker'lar tugashini kutamiz
	go func() {
		wg.Wait()
		close(results)
	}()

	// 5. Results channel'dan natijalarni o'qiymiz
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Println("Result:", result)
	}

	fmt.Println("All jobs completed!")

}
