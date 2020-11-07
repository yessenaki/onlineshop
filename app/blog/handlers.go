package blog

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/yesseneon/onlineshop/admin/post"
	"github.com/yesseneon/onlineshop/admin/post/category"
	"github.com/yesseneon/onlineshop/admin/post/tag"
	"github.com/yesseneon/onlineshop/helper"
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
			var ctgID int
			var tagID int
			if r.FormValue("ctg") != "" || r.FormValue("tag") != "" {
				ctgID, _ = strconv.Atoi(r.FormValue("ctg"))
				tagID, _ = strconv.Atoi(r.FormValue("tag"))
				if ctgID <= 0 && tagID <= 0 {
					http.Error(w, http.StatusText(404), http.StatusNotFound)
					return
				}
			}

			posts, err := post.FindWithLimit(1, ctgID, tagID)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			data := struct {
				Header Header
				Posts  []post.Post
				CtgID  int
				TagID  int
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
					Link:    "blog",
				},
				Posts: posts,
				CtgID: ctgID,
				TagID: tagID,
			}

			helper.Render(w, "blog.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			load, err := strconv.Atoi(r.FormValue("load"))
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			ctgID, _ := strconv.Atoi(r.FormValue("ctgID"))
			tagID, _ := strconv.Atoi(r.FormValue("tagID"))

			posts, err := post.FindWithLimit(load, ctgID, tagID)
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

func Details() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/post/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			id, _ := strconv.Atoi(r.FormValue("id"))
			if id <= 0 {
				http.Error(w, http.StatusText(404), http.StatusNotFound)
				return
			}

			p, err := post.FindOne(id)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			postTags, err := post.FindTags(p.ID)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			ctgs, err := post.FindCategories()
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			tags, err := tag.FindAll()
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			data := struct {
				Header     Header
				Post       post.Post
				PostTags   []tag.Tag
				Categories []category.Category
				Tags       []tag.Tag
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
					Link:    "blog",
				},
				Post:       p,
				PostTags:   postTags,
				Categories: ctgs,
				Tags:       tags,
			}

			helper.Render(w, "blog_details.gohtml", data)
			return
		}

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	})
}
