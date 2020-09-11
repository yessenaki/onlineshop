package main

import (
	"net/http"
	"onlineshop/app/user"
)

func basic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := user.SessionExists(r)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		if result {
			cookie, _ := r.Cookie("session_id")
			cookie.MaxAge = 15
			http.SetCookie(w, cookie)
		}

		next.ServeHTTP(w, r)
	})
}
