package main

import (
	"net/http"
	"time"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	w.Write([]byte("Salom! Hozirgi vaqt: " + t))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/time", Greeting)
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
