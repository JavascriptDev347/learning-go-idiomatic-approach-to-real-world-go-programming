package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeoutMS int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1\. So'rovning mavjud kontekstidan bola timeout kontekst yaratamiz
			ctx, cancel := context.WithTimeout(r.Context(), time.Duration(timeoutMS)*time.Millisecond)
			defer cancel()

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("Amal muvaffaqiyatli yakunlandi!"))
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Request canceled, because:", err)
		http.Error(w, "So'rov vaqti tugadi (Timeout): "+err.Error(), http.StatusGatewayTimeout)
	}
}

func main() {
	fmt.Println("Exercise 1")
	mux := http.NewServeMux()
	mux.HandleFunc("/slow", SlowHandler)

	// Middleware'ni ulaymiz: timeoutni atigi 500 millisekund (0.5 soniya) qilib belgilaymiz
	timeoutMiddleware := TimeoutMiddleware(500)
	handlerWithMiddleware := timeoutMiddleware(mux)
	server := &http.Server{
		Addr:        ":8080",
		Handler:     handlerWithMiddleware,
		IdleTimeout: 10 * time.Second,
	}
	fmt.Println("Server :8080 portida ishga tushdi...")
	server.ListenAndServe()
}
