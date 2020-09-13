package home

import (
	"io"
	"net/http"
	"onlineshop/app/user"
	"onlineshop/config"
	"onlineshop/helper"
)

type CtxData struct {
	AuthUser user.User
	Data     interface{}
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxData := CtxData{
			AuthUser: helper.AuthUserFromContext(r.Context()),
		}

		if r.Method == http.MethodGet {
			err := config.Tpl.ExecuteTemplate(w, "home.gohtml", ctxData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
