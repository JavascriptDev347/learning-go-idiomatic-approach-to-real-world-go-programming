package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("HTTP so'rov keldi",
				"method", r.Method,
				"path", r.URL.Path,
				"ip", r.RemoteAddr,
				"duration", time.Since(start).String())
		})
	}
}
