package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	CurrencyURL = "https://open.er-api.com/v6/latest/USD"
	FactURL     = "https://uselessfacts.jsph.pl/api/v2/facts/random"
	JokeURL     = "https://v2.jokeapi.dev/joke/Any?type=twopart"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	results := make(chan Result, 3)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		currency, err := fetchCurrency(ctx)
		results <- Result{Currency: currency, Err: err}
	}()

	go func() {
		defer wg.Done()
		fact, err := fetchFact(ctx)
		results <- Result{Fact: fact, Err: err}
	}()

	go func() {
		defer wg.Done()
		joke, err := fetchJoke(ctx)
		results <- Result{Joke: joke, Err: err}
	}()

	// Barcha worker tugagach channel yopiladi.
	go func() {
		wg.Wait()
		close(results)
	}()

	timeout := time.After(8 * time.Second)

	// x, err := context.WithTimeout(context.Background(), 8*time.Second)

	received := 0

	for received < 3 {
		select {
		case result, ok := <-results:
			if !ok {
				fmt.Println("Channel closed")
				return
			}

			received++

			if result.Err != nil {
				fmt.Println("Error:", result.Err)
				continue
			}

			switch {
			case result.Currency != nil:
				fmt.Println("Currency:", result.Currency.Rates["KZT"])

			case result.Fact != nil:
				fmt.Println("Fact:", result.Fact.Text)

			case result.Joke != nil:
				fmt.Println("Joke:", result.Joke.Setup)
				fmt.Println("Answer:", result.Joke.Delivery)
			}

		case <-timeout:
			fmt.Println("Timeout! API response took too long.")
			return
		}
	}

	fmt.Println("All API requests completed!")
	// for i := 0; i < 3; i++ {
	// 	result := <-results
	// 	if result.Err != nil {
	// 		fmt.Println("Error:", result.Err)
	// 		continue
	// 	}

	// 	switch {
	// 	case result.Currency != nil:
	// 		fmt.Println("Currency:", result.Currency.Rates["KZT"])
	// 	case result.Fact != nil:
	// 		fmt.Println("Fact:", result.Fact.Text)
	// 	case result.Joke != nil:
	// 		fmt.Println("Joke:", result.Joke.Setup)
	// 	}
	// }

}

func fetchCurrency(ctx context.Context) (*Currency, error) {
	var currency Currency
	if err := fetch(ctx, CurrencyURL, &currency); err != nil {
		return nil, err
	}
	return &currency, nil
}

func fetchFact(ctx context.Context) (*Fact, error) {
	var fact Fact
	if err := fetch(ctx, FactURL, &fact); err != nil {
		return nil, err
	}
	return &fact, nil
}

func fetchJoke(ctx context.Context) (*Joke, error) {
	var joke Joke
	if err := fetch(ctx, JokeURL, &joke); err != nil {
		return nil, err
	}
	return &joke, nil
}
