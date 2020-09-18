package category

import (
	"net/http"
	"onlineshop/app/user"
	"onlineshop/config"
	"onlineshop/helper"
	"strconv"
	"strings"
)

func Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := helper.AuthUserFromContext(r.Context())

		action := defineAction(r)
		switch action {
		case "index":
			index(w, r, auth)
		case "create":
			create(w, r, auth)
		case "store":
			store(w, r)
		case "edit":
			edit(w, r, auth)
		case "update":
			update(w, r)
		case "destroy":
			destroy(w, r)
		case "notFound":
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}

func index(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth       user.User
		Categories []Category
	}

	ctgs, err := allCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Categories: ctgs,
	}
	err = config.Tpl.ExecuteTemplate(w, "category.gohtml", data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth       user.User
		Categories []Category
		Category   Category
	}

	ctgs, err := allCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Categories: ctgs,
	}
	err = config.Tpl.ExecuteTemplate(w, "category_form.gohtml", data)
	if err != nil {
		// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func store(w http.ResponseWriter, r *http.Request) {
	parentID, err := strconv.Atoi(r.FormValue("parent_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	ctg := &Category{
		Title:    r.FormValue("title"),
		ParentID: parentID,
	}
	_, err = ctg.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth       user.User
		Categories []Category
		Category   Category
	}

	id, err := strconv.Atoi(r.URL.Query()["id"][0])
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	ctg, err := oneCategory(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if ctg.ID < 1 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	ctgs, err := allCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Categories: ctgs,
		Category:   ctg,
	}
	err = config.Tpl.ExecuteTemplate(w, "category_form.gohtml", data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	parentID, err := strconv.Atoi(r.FormValue("parent_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	ctg := &Category{
		ID:       id,
		Title:    r.FormValue("title"),
		ParentID: parentID,
	}
	err = ctg.update()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	ctg := &Category{ID: id}
	err = ctg.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories/", http.StatusSeeOther)
	return
}

func defineAction(r *http.Request) string {
	p := strings.Trim(r.URL.Path, "/")
	if len(strings.Split(p, "/")) > 2 {
		return "notFound"
	}

	action := "notAllowed"
	switch r.Method {
	case "GET":
		idExists := false
		if id, ok := r.URL.Query()["id"]; ok {
			if _, err := strconv.Atoi(id[0]); err == nil {
				idExists = true
			}
		}
		queryAction, ok := r.URL.Query()["action"]

		if ok && queryAction[0] == "edit" && idExists {
			action = "edit"
		} else if ok && queryAction[0] == "create" {
			action = "create"
		} else {
			action = "index"
		}
	case "POST":
		action = "store"
	case "PUT":
		action = "update"
	case "DELETE":
		action = "destroy"
	}

	return action
}
