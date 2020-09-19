package size

import (
	"net/http"
	"onlineshop/app/user"
	"onlineshop/config"
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
		Auth  user.User
		Sizes []Size
	}

	sizes, err := AllSizes()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:  auth,
		Sizes: sizes,
	}
	err = config.Tpl.ExecuteTemplate(w, "size.gohtml", data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth user.User
		Size Size
	}

	data := Data{Auth: auth}
	err := config.Tpl.ExecuteTemplate(w, "size_form.gohtml", data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func store(w http.ResponseWriter, r *http.Request) {
	t, err := strconv.Atoi(r.FormValue("type"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	size := &Size{
		Size: r.FormValue("size"),
		Type: t,
	}

	_, err = size.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/sizes/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth user.User
		Size Size
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
		Auth: auth,
		Size: size,
	}
	err = config.Tpl.ExecuteTemplate(w, "size_form.gohtml", data)
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
