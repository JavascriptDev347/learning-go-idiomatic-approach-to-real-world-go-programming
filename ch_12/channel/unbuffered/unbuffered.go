package main

func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)

	go sender(ch)

	for i := 0; i < 5; i++ {
		value := <-ch
		println("Received:", value)
	}

	close(ch)

}
