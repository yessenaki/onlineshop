package home

import (
	"io"
	"net/http"
	"onlineshop/helper"
)

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx interface{}
		var path = map[string]string{
			"folder": "home",
			"file":   "index.gohtml",
		}

		if r.Method == http.MethodGet {
			helper.RenderTemplate(w, path, ctx)
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, "405 method not allowed", 405)
			return
		}
	})
}
