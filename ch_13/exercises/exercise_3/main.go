package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type TimeResponse struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func buildTimeResponse(t time.Time) TimeResponse {
	return TimeResponse{
		DayOfWeek:  t.Weekday().String(),
		DayOfMonth: t.Day(),
		Month:      t.Month().String(),
		Year:       t.Year(), Hour: t.Hour(),
		Minute: t.Minute(),
		Second: t.Second()}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		resp := buildTimeResponse(now)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "JSON yaratishda xatolik yuz berdi", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(now.Format(time.RFC3339)))
}

func loggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("incoming request", "method", r.Method, "url", r.URL.Path, "ip", r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	}
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logMD := loggingMiddleware(logger)

	mux := http.NewServeMux()

	mux.Handle("/time", logMD(http.HandlerFunc(timeHandler)))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux}

	logger.Info("Server ishga tushdi", "addr", s.Addr)
	err := s.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}
