package clothessize

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
		Auth         user.User
		Clothessizes []Clothessize
	}

	cths, err := allClothessizes()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:         auth,
		Clothessizes: cths,
	}
	err = config.Tpl.ExecuteTemplate(w, "clothessize.gohtml", data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth        user.User
		Clothessize Clothessize
	}

	data := Data{Auth: auth}
	err := config.Tpl.ExecuteTemplate(w, "clothessize_form.gohtml", data)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func store(w http.ResponseWriter, r *http.Request) {
	cth := &Clothessize{
		Size: r.FormValue("size"),
	}
	_, err := cth.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/clothes-sizes/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth        user.User
		Clothessize Clothessize
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	cth, err := oneClothessize(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if cth.ID < 1 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	data := Data{
		Auth:        auth,
		Clothessize: cth,
	}
	err = config.Tpl.ExecuteTemplate(w, "clothessize_form.gohtml", data)
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

	cth := &Clothessize{
		ID:   id,
		Size: r.FormValue("size"),
	}
	err = cth.update()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/clothes-sizes/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	cth := &Clothessize{ID: id}
	err = cth.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/clothes-sizes/", http.StatusSeeOther)
	return
}
