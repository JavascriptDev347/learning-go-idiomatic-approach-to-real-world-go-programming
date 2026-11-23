package tracker

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"os"
)

const HeaderName = "X-Request-ID"

type requestIDKey struct{} // unexported kalit tipi

func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

func RequestIDFromContext(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey{}).(string)
	return requestID, ok
}

// Middleware — kelgan so'rovda ID bo'lsa davom ettiradi, bo'lmasa yangisini yaratadi.

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(HeaderName)
		if id == "" {
			id = newID()
		}
		ctx := ContextWithRequestID(r.Context(), id) // r.Context()'dan boshlaymiz
		w.Header().Set(HeaderName, id)               // javobga ID qo'shamiz
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Request — tashqi so'rovga ctx'dagi request ID'ni header qilib qo'shadi.
// Shu tufayli bitta foydalanuvchi so'rovi bir nechta servis loglarida bitta ID bilan ko'rinadi.
func Request(req *http.Request) *http.Request {
	if id, ok := RequestIDFromContext(req.Context()); ok {
		req.Header.Set(HeaderName, id)
	}
	return req
}

var baseLogger = slog.New(slog.NewTextHandler(os.Stdout, nil))

// Log — ctx'da request ID bo'lsa, uni logger'ga qo'shib qaytaradi.
func Log(ctx context.Context) *slog.Logger {
	if id, ok := RequestIDFromContext(ctx); ok {
		return baseLogger.With("request_id", id)
	}
	return baseLogger
}

func newID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}
