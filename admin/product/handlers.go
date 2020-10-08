package product

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"onlineshop/admin/brand"
	"onlineshop/admin/category"
	"onlineshop/admin/color"
	"onlineshop/admin/filestg"
	"onlineshop/admin/size"
	"onlineshop/app/user"
	"onlineshop/helper"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		Auth     user.User
		Products []Product
	}

	prods, err := FindAll()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:     auth,
		Products: prods,
	}
	helper.Render(w, "product.gohtml", data)
	return
}

func create(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth       user.User
		Product    Product
		Categories []category.Category
		Brands     []brand.Brand
		Colors     []color.Color
		Sizes      []size.Size
	}

	ctgs, err := category.FindChilds()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	brands, err := brand.AllBrands()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	colors, err := color.AllColors()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	sizes, err := size.AllSizes()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Categories: ctgs,
		Brands:     brands,
		Colors:     colors,
		Sizes:      sizes,
	}
	helper.Render(w, "product_form.gohtml", data)
	return
}

func store(w http.ResponseWriter, r *http.Request, auth user.User) {
	price := priceToInt(r.FormValue("price"))
	gender, _ := strconv.Atoi(r.FormValue("gender"))

	var isKids int
	if r.FormValue("is_kids") == "on" {
		isKids = 1
	}

	var isNew int
	if r.FormValue("is_new") == "on" {
		isNew = 1
	}

	brandID, _ := strconv.Atoi(r.FormValue("brand_id"))
	colorID, _ := strconv.Atoi(r.FormValue("color_id"))
	ctgID, _ := strconv.Atoi(r.FormValue("category_id"))
	sizeID, _ := strconv.Atoi(r.FormValue("size_id"))

	var isDiscount int
	var dscPercent int
	var oldPrice int
	if r.FormValue("is_discount") == "on" {
		isDiscount = 1
		dscPercent, _ = strconv.Atoi(r.FormValue("dsc_percent"))
		oldPrice = priceToInt(r.FormValue("old_price"))
	}

	prod := &Product{
		Title:       r.FormValue("title"),
		Price:       price,
		OldPrice:    oldPrice,
		Gender:      gender,
		IsKids:      isKids,
		IsNew:       isNew,
		IsDiscount:  isDiscount,
		DscPercent:  dscPercent,
		BrandID:     brandID,
		ColorID:     colorID,
		CategoryID:  ctgID,
		SizeID:      sizeID,
		Description: r.FormValue("description"),
	}

	if prod.validate(r) == false {
		type Data struct {
			Auth       user.User
			Product    *Product
			Categories []category.Category
			Brands     []brand.Brand
			Colors     []color.Color
			Sizes      []size.Size
		}

		ctgs, err := category.AllCategories()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		brands, err := brand.AllBrands()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		colors, err := color.AllColors()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		sizes, err := size.AllSizes()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := Data{
			Auth:       auth,
			Product:    prod,
			Categories: ctgs,
			Brands:     brands,
			Colors:     colors,
			Sizes:      sizes,
		}
		helper.Render(w, "product_form.gohtml", data)
		return
	}

	id, err := prod.store()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = uploadImages(id, r)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/products/", http.StatusSeeOther)
	return
}

func edit(w http.ResponseWriter, r *http.Request, auth user.User) {
	type Data struct {
		Auth       user.User
		Product    Product
		Categories []category.Category
		Brands     []brand.Brand
		Colors     []color.Color
		Sizes      []size.Size
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	prod, err := FindOne(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	ctgs, err := category.FindChilds()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	brands, err := brand.AllBrands()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	colors, err := color.AllColors()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	sizes, err := size.AllSizes()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	data := Data{
		Auth:       auth,
		Product:    prod,
		Categories: ctgs,
		Brands:     brands,
		Colors:     colors,
		Sizes:      sizes,
	}
	helper.Render(w, "product_form.gohtml", data)
	return
}

func update(w http.ResponseWriter, r *http.Request, auth user.User) {
	price := priceToInt(r.FormValue("price"))
	gender, _ := strconv.Atoi(r.FormValue("gender"))

	var isKids int
	if r.FormValue("is_kids") == "on" {
		isKids = 1
	}

	var isNew int
	if r.FormValue("is_new") == "on" {
		isNew = 1
	}

	brandID, _ := strconv.Atoi(r.FormValue("brand_id"))
	colorID, _ := strconv.Atoi(r.FormValue("color_id"))
	ctgID, _ := strconv.Atoi(r.FormValue("category_id"))
	sizeID, _ := strconv.Atoi(r.FormValue("size_id"))

	var isDiscount int
	var dscPercent int
	var oldPrice int
	if r.FormValue("is_discount") == "on" {
		isDiscount = 1
		dscPercent, _ = strconv.Atoi(r.FormValue("dsc_percent"))
		oldPrice = priceToInt(r.FormValue("old_price"))
	}

	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prod := &Product{
		ID:          id,
		Title:       r.FormValue("title"),
		Price:       price,
		OldPrice:    oldPrice,
		Gender:      gender,
		IsKids:      isKids,
		IsNew:       isNew,
		IsDiscount:  isDiscount,
		DscPercent:  dscPercent,
		BrandID:     brandID,
		ColorID:     colorID,
		CategoryID:  ctgID,
		SizeID:      sizeID,
		Description: r.FormValue("description"),
	}

	if prod.validate(r) == false {
		type Data struct {
			Auth       user.User
			Product    *Product
			Categories []category.Category
			Brands     []brand.Brand
			Colors     []color.Color
			Sizes      []size.Size
		}

		ctgs, err := category.AllCategories()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		brands, err := brand.AllBrands()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		colors, err := color.AllColors()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		sizes, err := size.AllSizes()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		data := Data{
			Auth:       auth,
			Product:    prod,
			Categories: ctgs,
			Brands:     brands,
			Colors:     colors,
			Sizes:      sizes,
		}
		helper.Render(w, "product_form.gohtml", data)
		return
	}

	err = prod.update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/products/", http.StatusSeeOther)
	return
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("_id"))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	prod := &Product{ID: id}
	err = prod.destroy()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/products/", http.StatusSeeOther)
	return
}

func priceToInt(price string) int {
	fprice, _ := strconv.ParseFloat(price, 64)
	rprice := math.Round(fprice * 100)

	return int(rprice)
}

func uploadImages(id int, r *http.Request) error {
	// define some variables used throughout the function
	// n: for keeping track of bytes read and written
	// err: for storing errors that need checking
	var n int
	var err error

	// define pointers for the multipart reader and its parts
	var mr *multipart.Reader
	var part *multipart.Part

	if mr, err = r.MultipartReader(); err != nil {
		log.Printf("Hit error while opening multipart reader: %s", err.Error())
		return err
	}

	// buffer to be used for reading bytes from files
	chunk := make([]byte, 4096)

	// continue looping through all parts, *multipart.Reader.NextPart() will
	// return an End of File when all parts have been read.
	for {
		// variables used in this loop only
		// tempfile: filehandler for the temporary file
		// filesize: how many bytes where written to the tempfile
		// uploaded: boolean to flip when the end of a part is reached
		var tempfile *os.File
		var filesize int
		var uploaded bool

		if part, err = mr.NextPart(); err != nil {
			if err != io.EOF {
				log.Printf("Hit error while fetching next part: %s", err.Error())
				return err
			}

			log.Printf("Hit last part of multipart upload")
			return nil
		}
		// at this point the filename and the mimetype is known
		// filename: part.FileName()
		// mimetype: part.Header
		ext := strings.Split(part.FileName(), ".")[1]

		tempfile, err = ioutil.TempFile(os.TempDir(), "upload-*.tmp")
		if err != nil {
			return err
		}
		defer tempfile.Close()

		// defer the removal of the tempfile as well, something can be done
		// with it before the function is over (as long as you have the filehandle)
		defer os.Remove(tempfile.Name())

		// continue reading until the whole file is upload or an error is reached
		for !uploaded {
			if n, err = part.Read(chunk); err != nil {
				if err != io.EOF {
					log.Printf("Hit error while reading chunk: %s", err.Error())
					return err
				}
				uploaded = true
			}

			if n, err = tempfile.Write(chunk[:n]); err != nil {
				log.Printf("Hit error while writing chunk: %s", err.Error())
				return err
			}
			filesize += n
		}

		// once uploaded something can be done with the file, the last defer
		// statement will remove the file after the function returns so any
		// errors during upload won't hit this, but at least the tempfile is
		// cleaned up
		if n, err := tempfile.Seek(0, 0); err != nil || n != 0 {
			log.Printf("unable to seek to beginning of file '%s'", tempfile.Name())
		}

		h := sha256.New()
		if _, err := io.Copy(h, tempfile); err != nil {
			log.Printf("unable to hash '%s': %s", tempfile.Name(), err.Error())
		}
		filename := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// get working directory
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		path := filepath.Join(wd, "static", "uploads", filename)

		newFile, err := os.Create(path)
		if err != nil {
			log.Println(err)
		}
		defer newFile.Close()

		if n, err := tempfile.Seek(0, 0); err != nil || n != 0 {
			log.Printf("unable to seek to beginning of file '%s'", tempfile.Name())
		}

		if _, err := io.Copy(newFile, tempfile); err != nil {
			log.Println(err)
		}

		fs := filestg.Filestg{
			Name:      part.FileName(),
			Path:      "/static/uploads/" + filename,
			ProductID: id,
		}
		_, err = fs.Store()
		return err
	}
}
