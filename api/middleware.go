package api

import (
	"net/http"
)

type KeyAuthMiddle struct {
	key    string
	secret string
}

func AuthMiddleware(key, secret string) *KeyAuthMiddle {
	return &KeyAuthMiddle{
		key:    key,
		secret: secret,
	}
}

func (m *KeyAuthMiddle) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			key    = r.Header.Get("x-auth-key")
			secret = r.Header.Get("x-auth-secret")
		)
		if key != m.key || secret != m.secret {
			w.WriteHeader(400)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
