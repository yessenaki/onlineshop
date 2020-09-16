package category

import (
	"io"
	"net/http"
	"onlineshop/app/user"
	"onlineshop/config"
	"onlineshop/helper"
	"strconv"
)

type CtxData struct {
	AuthUser   user.User
	Categories []Category
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxData := CtxData{
			AuthUser: helper.AuthUserFromContext(r.Context()),
		}

		if r.Method == http.MethodGet {
			ctgs, err := allCategories()
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			ctxData.Categories = ctgs
			err = config.Tpl.ExecuteTemplate(w, "category.gohtml", ctxData)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			return
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}

func Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxData := CtxData{
			AuthUser: helper.AuthUserFromContext(r.Context()),
		}

		if r.Method == http.MethodGet {
			ctgs, err := allCategories()
			if err != nil {
				// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ctxData.Categories = ctgs
			err = config.Tpl.ExecuteTemplate(w, "category_create.gohtml", ctxData)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			return
		} else if r.Method == http.MethodPost {
			parentID, err := strconv.Atoi(r.PostFormValue("parent_id"))
			if err != nil {
				// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ctg := &Category{
				Title:    r.PostFormValue("title"),
				ParentID: parentID,
			}

			_, err = ctg.save()
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
			return
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
