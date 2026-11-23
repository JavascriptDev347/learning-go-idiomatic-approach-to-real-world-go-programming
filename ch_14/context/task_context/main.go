package main

import (
	"log/slog"
	"net/http"
	"os"
	"task_context/internal/identity"
	"task_context/internal/order"
	"task_context/internal/stock"
	"task_context/internal/tracker"
	"time"
)

func main() {
	ctrl := order.Controller{
		Service: order.Service{
			Client:   &http.Client{Timeout: 5 * time.Second},
			StockURL: "http://localhost:8080/internal/stock/",
		},
	}

	mux := http.NewServeMux() // DefaultServeMux emas
	mux.Handle("GET /orders/{id}", identity.Middleware(http.HandlerFunc(ctrl.GetOrder)))
	mux.HandleFunc("GET /internal/stock/{id}", stock.Handler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      tracker.Middleware(mux), // hamma so'rov request_id oladi
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	slog.Info("listening", "addr", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}
