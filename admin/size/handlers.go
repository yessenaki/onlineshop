package size

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
		Context helper.ContextData
		Sizes   []Size
	}

	sizes, err := AllSizes()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Context: ctx,
		Sizes:   sizes,
	}
	helper.Render(w, "size.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context helper.ContextData
		Size    Size
	}

	data := Data{
		Context: ctx,
	}
	helper.Render(w, "size_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	t, err := strconv.Atoi(r.FormValue("type"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	size := &Size{
		Size: r.FormValue("size"),
		Type: t,
	}

	if size.validate() == false {
		type Data struct {
			Context helper.ContextData
			Size    *Size
		}

		data := Data{
			Context: ctx,
			Size:    size,
		}
		helper.Render(w, "size_form.gohtml", data)
		return
	}

	_, err = size.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/sizes/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context helper.ContextData
		Size    Size
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	size, err := oneSize(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if size.ID < 1 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	data := Data{
		Context: ctx,
		Size:    size,
	}
	helper.Render(w, "size_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	t, err := strconv.Atoi(r.FormValue("type"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	size := &Size{
		ID:   id,
		Size: r.FormValue("size"),
		Type: t,
	}

	if size.validate() == false {
		type Data struct {
			Context helper.ContextData
			Size    *Size
		}

		data := Data{
			Context: ctx,
			Size:    size,
		}
		helper.Render(w, "size_form.gohtml", data)
		return
	}

	err = size.update()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/sizes/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	size := &Size{ID: id}
	err = size.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/sizes/", http.StatusSeeOther)
	return
}
