package identity

import (
	"context"
	"net/http"
	"task_context/internal/tracker"
)

type userKey struct{}

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userKey{}).(string)
	return user, ok
}

// Middleware — header'dan user'ni o'qib ctx'ga qo'yadi.
// Soddalik uchun X-User header; real hayotda cookie yoki JWT.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-User")
		if user == "" {
			tracker.Log(r.Context()).Warn("unauthorized request")
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := ContextWithUser(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx)) // yangi ctx'ni so'rovga biriktiramiz!
	})
}
