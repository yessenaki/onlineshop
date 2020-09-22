package product

import (
	"math"
	"net/http"
	"onlineshop/admin/brand"
	"onlineshop/admin/category"
	"onlineshop/admin/color"
	"onlineshop/admin/size"
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

	prods, err := allProducts()
	if err != nil {
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
		Title:      r.FormValue("title"),
		Price:      price,
		OldPrice:   oldPrice,
		Gender:     gender,
		IsKids:     isKids,
		IsNew:      isNew,
		IsDiscount: isDiscount,
		DscPercent: dscPercent,
		BrandID:    brandID,
		ColorID:    colorID,
		CategoryID: ctgID,
		SizeID:     sizeID,
	}

	if prod.validate() == false {
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

	_, err := prod.store()
	if err != nil {
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

	prod, err := oneProduct(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
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
	err = config.Tpl.ExecuteTemplate(w, "product_form.gohtml", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
		ID:         id,
		Title:      r.FormValue("title"),
		Price:      price,
		OldPrice:   oldPrice,
		Gender:     gender,
		IsKids:     isKids,
		IsNew:      isNew,
		IsDiscount: isDiscount,
		DscPercent: dscPercent,
		BrandID:    brandID,
		ColorID:    colorID,
		CategoryID: ctgID,
		SizeID:     sizeID,
	}

	if prod.validate() == false {
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
