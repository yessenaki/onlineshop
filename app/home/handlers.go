package home

import (
	"io"
	"net/http"
	"onlineshop/app/user"
	"onlineshop/config"
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

		data := struct {
			Header Header
		}{
			Header: Header{
				Auth: helper.AuthUserFromContext(r.Context()),
				Link: "home",
			},
		}

		if r.Method == http.MethodGet {
			err := config.Tpl.ExecuteTemplate(w, "home.gohtml", data)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
