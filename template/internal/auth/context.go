package auth

import (
	"context"
	"net/http"
)

type contextKey string

var userContextKey contextKey = "user"

func UserFromContext(ctx context.Context) *User {
	if user, ok := ctx.Value(userContextKey).(*User); ok {
		return user
	}
	return nil
}

func WithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func Middleware(local bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if local {
				r = r.WithContext(WithUser(r.Context(), &LocalUser))
			} else {
				// get user from token
				// TODO: implement token parsing
			}
			next.ServeHTTP(w, r)
		})
	}
}
