package home

import (
	"io"
	"net/http"
	"onlineshop/app/user"
	"onlineshop/helper"
)

// Header struct
type Header struct {
	Auth user.User
	Link string
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			data := struct {
				Header Header
			}{
				Header: Header{
					Auth: r.Context().Value(helper.AuthUserKey).(user.User),
					Link: "home",
				},
			}

			helper.Render(w, "home.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
