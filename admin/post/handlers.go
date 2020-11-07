package post

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/yesseneon/onlineshop/admin/post/category"
	"github.com/yesseneon/onlineshop/admin/post/tag"
	"github.com/yesseneon/onlineshop/helper"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	posts, err := FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := struct {
		Context helper.ContextData
		Posts   []Post
	}{
		Context: ctx,
		Posts:   posts,
	}
	helper.Render(w, "post.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	ctgs, err := category.FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tags, err := tag.FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := struct {
		Context    helper.ContextData
		Categories []category.Category
		Tags       []tag.Tag
		Post       Post
	}{
		Context:    ctx,
		Categories: ctgs,
		Tags:       tags,
	}
	helper.Render(w, "post_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	r.ParseForm()

	ctgID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	post := &Post{
		Title:      r.FormValue("title"),
		Body:       r.FormValue("body"),
		CategoryID: ctgID,
		Tags:       r.Form["tags"],
		Author:     r.FormValue("author"),
	}

	success, err := post.validate(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if success == false {
		ctgs, err := category.FindAll()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		tags, err := tag.FindAll()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := struct {
			Context    helper.ContextData
			Categories []category.Category
			Tags       []tag.Tag
			Post       *Post
		}{
			Context:    ctx,
			Categories: ctgs,
			Tags:       tags,
			Post:       post,
		}
		helper.Render(w, "post_form.gohtml", data)
		return
	}

	m, err := uploadImage(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	post.ImagePath = m["ImagePath"]
	post.ImageName = m["ImageName"]

	_, err = post.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/posts/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	if id <= 0 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	ctgs, err := category.FindAll()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tags, err := tag.FindWithSelected(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	post, err := FindOne(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := struct {
		Context    helper.ContextData
		Categories []category.Category
		Tags       []tag.Tag
		Post       Post
	}{
		Context:    ctx,
		Categories: ctgs,
		Tags:       tags,
		Post:       post,
	}
	helper.Render(w, "post_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, ctx helper.ContextData) {
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("_id"))
	ctgID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	post := &Post{
		ID:         id,
		Title:      r.FormValue("title"),
		Body:       r.FormValue("body"),
		ImagePath:  r.FormValue("image_path"),
		ImageName:  r.FormValue("image_name"),
		CategoryID: ctgID,
		Tags:       r.Form["tags"],
		Author:     r.FormValue("author"),
	}

	success, err := post.validate(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if success == false {
		ctgs, err := category.FindAll()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		tags, err := tag.FindAll()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := struct {
			Context    helper.ContextData
			Categories []category.Category
			Tags       []tag.Tag
			Post       *Post
		}{
			Context:    ctx,
			Categories: ctgs,
			Tags:       tags,
			Post:       post,
		}
		helper.Render(w, "post_form.gohtml", data)
		return
	}

	m, err := uploadImage(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if m != nil {
		post.ImagePath = m["ImagePath"]
		post.ImageName = m["ImageName"]
	}

	err = post.update()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/posts/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("_id"))
	post := &Post{ID: id}
	err := post.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/posts/", http.StatusSeeOther)
	return
}

func uploadImage(r *http.Request) (map[string]string, error) {
	f, fh, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	ext := strings.Split(fh.Filename, ".")[1]
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Printf("Unable to hash '%s': %s", fh.Filename, err.Error())
	}
	filename := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

	wd, err := os.Getwd() // working directory
	if err != nil {
		return nil, err
	}
	path := filepath.Join(wd, "static", "uploads", filename)

	nf, err := os.Create(path) // new file
	if err != nil {
		return nil, err
	}
	defer nf.Close()

	if n, err := f.Seek(0, 0); err != nil || n != 0 {
		log.Printf("Unable to seek to beginning of file '%s'", filename)
	}
	if _, err := io.Copy(nf, f); err != nil {
		log.Printf("Unable to copy '%s': %s", filename, err.Error())
	}

	m := map[string]string{
		"ImagePath": "/assets/uploads/" + filename,
		"ImageName": fh.Filename,
	}

	return m, nil
}
