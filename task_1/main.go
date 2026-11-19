package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"task_1/middleware"
	"task_1/model"
	"task_1/service"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	service := service.NewAggregatorService(logger)

	mux := http.NewServeMux()

	// get aggregated posts
	mux.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		posts, _, err := service.AggregatePosts(20, 4)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(posts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /stats", func(w http.ResponseWriter, R *http.Request) {
		posts, dur, err := service.AggregatePosts(20, 4)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		stats := model.StatsResponse{TotalProcessed: len(posts), Duration: dur.String(), DurationNano: dur}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(stats)
	})

	// middlewares
	loggingMiddleware := middleware.RequestLogger(logger)
	handlerWithMiddleware := loggingMiddleware(mux)

	// server

	s := http.Server{
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handlerWithMiddleware,
	}

	logger.Info("PostAggregator serveri ishga tushdi", "port", 8080)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Server to'xtadi", "err", err)
		os.Exit(1)
	}
}
