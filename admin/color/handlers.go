package color

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
		Context helper.ContextData
		Colors  []Color
	}

	colors, err := AllColors()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Context: ctx,
		Colors:  colors,
	}
	helper.Render(w, "color.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context helper.ContextData
		Color   Color
	}

	data := Data{
		Context: ctx,
	}
	helper.Render(w, "color_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	color := &Color{
		Name: r.FormValue("name"),
	}

	if color.validate() == false {
		type Data struct {
			Context helper.ContextData
			Color   *Color
		}

		data := Data{
			Context: ctx,
			Color:   color,
		}
		helper.Render(w, "color_form.gohtml", data)
		return
	}

	_, err := color.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/colors/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	type Data struct {
		Context helper.ContextData
		Color   Color
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	color, err := oneColor(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if color.ID < 1 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	data := Data{
		Context: ctx,
		Color:   color,
	}
	helper.Render(w, "color_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	color := &Color{
		ID:   id,
		Name: r.FormValue("name"),
	}

	if color.validate() == false {
		type Data struct {
			Context helper.ContextData
			Color   *Color
		}

		data := Data{
			Context: ctx,
			Color:   color,
		}
		helper.Render(w, "color_form.gohtml", data)
		return
	}

	err = color.update()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/colors/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	color := &Color{ID: id}
	err = color.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/colors/", http.StatusSeeOther)
	return
}
