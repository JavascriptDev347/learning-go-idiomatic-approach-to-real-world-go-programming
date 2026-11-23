package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
)

type levelKey struct{}

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, levelKey{}, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(levelKey{}).(Level)
	return level, ok
}

func Log(ctx context.Context, level Level, message string) {
	currentLevel, ok := LevelFromContext(ctx)
	if !ok {
		currentLevel = LevelInfo // Default level if not set
	}
	if currentLevel == LevelInfo && level == LevelDebug {
		return // Skip debug messages if current level is info
	}
	fmt.Printf("[%s] %s\n", strings.ToUpper(string(level)), message)
}

func LogLevelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		queryLevel := r.URL.Query().Get("log_level")
		if queryLevel == string(LevelDebug) {
			ctx = ContextWithLevel(ctx, LevelDebug)
		} else {
			ctx = ContextWithLevel(ctx, LevelInfo)
		}

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	Log(ctx, LevelInfo, "Foydalanuvchi sahifasiga so'rov keldi")
	Log(ctx, LevelDebug, "Maxfiy debug ma'lumoti: kesh tekshirilmoqda, user_id=42")
	w.Write([]byte("Salom, foydalanuvchi!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", UserHandler)
	wrappedMux := LogLevelMiddleware(mux)
	fmt.Println("Server :8080 portida ishga tushdi...")
	http.ListenAndServe(":8080", wrappedMux)
}
