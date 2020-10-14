package tag

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
	tags, err := FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := struct {
		Context helper.ContextData
		Tags    []Tag
	}{
		Context: ctx,
		Tags:    tags,
	}
	helper.Render(w, "post_tag.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	data := struct {
		Context helper.ContextData
		Tag     Tag
	}{
		Context: ctx,
	}
	helper.Render(w, "post_tag_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	tag := &Tag{
		Name: r.FormValue("name"),
	}

	if tag.validate() == false {
		data := struct {
			Context helper.ContextData
			Tag     *Tag
		}{
			Context: ctx,
			Tag:     tag,
		}
		helper.Render(w, "post_tag_form.gohtml", data)
		return
	}

	_, err := tag.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/post-tags/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	if id <= 0 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	tag, err := findOne(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := struct {
		Context helper.ContextData
		Tag     Tag
	}{
		Context: ctx,
		Tag:     tag,
	}
	helper.Render(w, "post_tag_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	id, _ := strconv.Atoi(r.FormValue("_id"))
	tag := &Tag{
		ID:   id,
		Name: r.FormValue("name"),
	}

	if tag.validate() == false {
		data := struct {
			Context helper.ContextData
			Tag     *Tag
		}{
			Context: ctx,
			Tag:     tag,
		}
		helper.Render(w, "post_tag_form.gohtml", data)
		return
	}

	err := tag.update()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/post-tags/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("_id"))
	tag := &Tag{ID: id}
	err := tag.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/post-tags/", http.StatusSeeOther)
	return
}
