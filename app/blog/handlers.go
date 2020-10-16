package blog

import (
	"encoding/json"
	"log"
	"net/http"
	"onlineshop/admin/post"
	"onlineshop/helper"
	"strconv"
)

// Header struct
type Header struct {
	Context helper.ContextData
	Link    string
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/blog/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			posts, err := post.FindWithLimit(1)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			data := struct {
				Header Header
				Posts  []post.Post
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
					Link:    "blog",
				},
				Posts: posts,
			}

			helper.Render(w, "blog.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			load, err := strconv.Atoi(r.FormValue("load"))
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			posts, err := post.FindWithLimit(load)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			data := struct {
				Posts []post.Post
			}{
				Posts: posts,
			}

			j, err := json.Marshal(data)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(j)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
