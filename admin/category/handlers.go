package category

import (
	"net/http"
	"github.com/yesseneon/onlineshop/helper"
	"strconv"
)

func Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := helper.GetContextData(r.Context())
		action := helper.DefineAction(r)
		switch action {
		case "index":
			index(w, r, ctx)
		case "create":
			create(w, r, ctx)
		case "store":
			store(w, r, ctx)
		case "edit":
			edit(w, r, ctx)
		case "update":
			update(w, r, ctx)
		case "destroy":
			destroy(w, r)
		case "notFound":
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		}
	})
}

func index(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context    helper.ContextData
		Categories []Category
	}

	ctgs, err := AllCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Context:    ctx,
		Categories: ctgs,
	}
	helper.Render(w, "category.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context    helper.ContextData
		Categories []Category
		Category   Category
	}

	ctgs, err := AllCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Context:    ctx,
		Categories: ctgs,
	}
	helper.Render(w, "category_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	parentID, _ := strconv.Atoi(r.FormValue("parent_id"))
	gender, _ := strconv.Atoi(r.FormValue("gender"))
	isKids, _ := strconv.Atoi(r.FormValue("is_kids"))

	ctg := &Category{
		Name:     r.FormValue("name"),
		ParentID: parentID,
		Gender:   gender,
		IsKids:   isKids,
	}

	if ctg.validate() == false {
		type Data struct {
			Context    helper.ContextData
			Categories []Category
			Category   *Category
		}

		ctgs, err := AllCategories()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := Data{
			Context:    ctx,
			Categories: ctgs,
			Category:   ctg,
		}
		helper.Render(w, "category_form.gohtml", data)
		return
	}

	_, err := ctg.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context    helper.ContextData
		Categories []Category
		Category   Category
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	ctg, err := FindOne(id)
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
		Context:    ctx,
		Categories: ctgs,
		Category:   ctg,
	}
	helper.Render(w, "category_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	id, _ := strconv.Atoi(r.FormValue("_id"))
	parentID, _ := strconv.Atoi(r.FormValue("parent_id"))
	gender, _ := strconv.Atoi(r.FormValue("gender"))
	isKids, _ := strconv.Atoi(r.FormValue("is_kids"))

	ctg := &Category{
		ID:       id,
		Name:     r.FormValue("name"),
		ParentID: parentID,
		Gender:   gender,
		IsKids:   isKids,
	}

	if ctg.validate() == false {
		type Data struct {
			Context    helper.ContextData
			Categories []Category
			Category   *Category
		}

		ctgs, err := AllCategories()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := Data{
			Context:    ctx,
			Categories: ctgs,
			Category:   ctg,
		}
		helper.Render(w, "category_form.gohtml", data)
		return
	}

	err := ctg.update()
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
