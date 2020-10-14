package category

import (
	"net/http"
	"onlineshop/helper"
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

	ctgs, err := FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Context:    ctx,
		Categories: ctgs,
	}
	helper.Render(w, "post_category.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context  helper.ContextData
		Category Category
	}

	data := Data{
		Context: ctx,
	}
	helper.Render(w, "post_category_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	ctg := &Category{
		Name: r.FormValue("name"),
	}

	if ctg.validate() == false {
		type Data struct {
			Context  helper.ContextData
			Category *Category
		}

		data := Data{
			Context:  ctx,
			Category: ctg,
		}
		helper.Render(w, "post_category_form.gohtml", data)
		return
	}

	_, err := ctg.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/post-categories/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context  helper.ContextData
		Category Category
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	if id <= 0 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	ctg, err := findOne(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Context:  ctx,
		Category: ctg,
	}
	helper.Render(w, "post_category_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	id, _ := strconv.Atoi(r.FormValue("_id"))
	ctg := &Category{
		ID:   id,
		Name: r.FormValue("name"),
	}

	if ctg.validate() == false {
		type Data struct {
			Context  helper.ContextData
			Category *Category
		}

		data := Data{
			Context:  ctx,
			Category: ctg,
		}
		helper.Render(w, "post_category_form.gohtml", data)
		return
	}

	err := ctg.update()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/post-categories/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("_id"))
	ctg := &Category{ID: id}
	err := ctg.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/post-categories/", http.StatusSeeOther)
	return
}
