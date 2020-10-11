package color

import (
	"net/http"
	"onlineshop/app/user"
	"onlineshop/helper"
	"strconv"
)

func Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Context().Value(helper.AuthUserKey).(user.User)
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
		Auth   user.User
		Colors []Color
	}

	colors, err := AllColors()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:   auth,
		Colors: colors,
	}
	helper.Render(w, "color.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth  user.User
		Color Color
	}

	data := Data{Auth: auth}
	helper.Render(w, "color_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, auth user.User) {
	color := &Color{
		Name: r.FormValue("name"),
	}

	if color.validate() == false {
		type Data struct {
			Auth  user.User
			Color *Color
		}

		data := Data{
			Auth:  auth,
			Color: color,
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

func edit(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth  user.User
		Color Color
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
		Auth:  auth,
		Color: color,
	}
	helper.Render(w, "color_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, auth user.User) {
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
			Auth  user.User
			Color *Color
		}

		data := Data{
			Auth:  auth,
			Color: color,
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
