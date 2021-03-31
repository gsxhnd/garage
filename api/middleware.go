package api

import "net/http"

type HTTPAuthMiddle struct {
	user     string
	password string
}

func AuthMdiddleware(username, password string) *HTTPAuthMiddle {
	return &HTTPAuthMiddle{
		user:     username,
		password: password,
	}
}

func (m *HTTPAuthMiddle) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUser, reqPasswd, hasAuth := r.BasicAuth()
		if (m.user == "" && m.password == "") ||
			(hasAuth && reqUser == m.user && reqPasswd == m.password) {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}
