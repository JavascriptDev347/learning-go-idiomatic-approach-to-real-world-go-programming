package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {

	mux := http.NewServeMux()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	md := LoggingMiddleware(logger)

	mux.Handle("/", md(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Salom! Hozirgi vaqt: " + time.Now().Format(time.RFC3339)))
	})))

	s := http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			method := r.Method
			url := r.URL.Path
			logger.Info("incoming request", "method", method, "url", url, "ipp address:", ip)
			next.ServeHTTP(w, r)
		})
	}
}
