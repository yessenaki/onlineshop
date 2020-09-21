package category

import (
	"net/http"
	"onlineshop/app/user"
	"onlineshop/helper"
	"strconv"
)

func Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := helper.AuthUserFromContext(r.Context())
		action := helper.DefineAction(r)
		switch action {
		case "index":
			index(w, r, auth)
		case "create":
			create(w, r, auth)
		case "store":
			store(w, r, auth)
		case "edit":
			edit(w, r, auth)
		case "update":
			update(w, r, auth)
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

	ctgs, err := AllCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Categories: ctgs,
	}
	helper.Render(w, "category.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth       user.User
		Categories []Category
		Category   Category
	}

	ctgs, err := AllCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Categories: ctgs,
	}
	helper.Render(w, "category_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, auth user.User) {
	parentID, err := strconv.Atoi(r.FormValue("parent_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	ctg := &Category{
		Name:     r.FormValue("name"),
		ParentID: parentID,
	}

	if ctg.validate() == false {
		type Data struct {
			Auth       user.User
			Categories []Category
			Category   *Category
		}

		ctgs, err := AllCategories()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := Data{
			Auth:       auth,
			Categories: ctgs,
			Category:   ctg,
		}
		helper.Render(w, "category_form.gohtml", data)
		return
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

	id, err := strconv.Atoi(r.FormValue("id"))
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

	ctgs, err := AllCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Categories: ctgs,
		Category:   ctg,
	}
	helper.Render(w, "category_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, auth user.User) {
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
		Name:     r.FormValue("name"),
		ParentID: parentID,
	}

	if ctg.validate() == false {
		type Data struct {
			Auth       user.User
			Categories []Category
			Category   *Category
		}

		ctgs, err := AllCategories()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := Data{
			Auth:       auth,
			Categories: ctgs,
			Category:   ctg,
		}
		helper.Render(w, "category_form.gohtml", data)
		return
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
