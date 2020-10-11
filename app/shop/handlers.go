package shop

import (
	"fmt"
	"math"
	"net/http"
	"onlineshop/admin/brand"
	"onlineshop/admin/category"
	"onlineshop/admin/color"
	"onlineshop/admin/file"
	"onlineshop/admin/product"
	"onlineshop/admin/size"
	"onlineshop/app/user"
	"onlineshop/helper"
	"strconv"
)

// Header struct
type Header struct {
	Auth user.User
	Link string
}

// Pagination struct
type Pagination struct {
	Bprev   int // before the previous
	Prev    int
	Current int
	Next    int
	Anext   int // after the next
}

// Qstr struct, Qstr = Query string
type Qstr struct {
	Stype  string
	Gender string
	Ctg    string
	Brands string
	Sizes  string
	Colors string
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/shop/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			stype := "women" // shop type
			gender := 1      // female
			isKids := 0      // not for kids

			if r.FormValue("t") == "men" {
				stype = "men"
				gender = 0
			} else if r.FormValue("t") == "kids" {
				stype = "kids"
				isKids = 1
				if r.FormValue("g") == "boy" {
					gender = 0
				}
			}

			ctgID, _ := strconv.Atoi(r.FormValue("ctg"))
			ctgs, err := category.ByParams(gender, isKids, ctgID)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			checkedBrands := helper.ListToSlice(r.FormValue("b"))
			brands, err := brand.Arrange(checkedBrands)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			checkedSizes := helper.ListToSlice(r.FormValue("s"))
			sizes, err := size.Arrange(checkedSizes, ctgID)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			checkedColors := helper.ListToSlice(r.FormValue("c"))
			colors, err := color.Arrange(checkedColors)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			page, _ := strconv.Atoi(r.FormValue("page"))
			if page == 0 {
				page = 1
			}
			params := map[string]interface{}{
				"gender": gender,
				"isKids": isKids,
				"ctgID":  ctgID,
				"brands": r.FormValue("b"),
				"sizes":  r.FormValue("s"),
				"colors": r.FormValue("c"),
				"page":   page,
			}
			prods, quantity, err := product.FindByParams(params)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			pgn := createPgn(page, quantity)
			qstr := arrangeQstr(r)
			qstr.Stype = "?t=" + stype

			data := struct {
				Header Header
				Ctgs   []category.Parent
				Brands []brand.Brand
				Sizes  []size.Size
				Colors []color.Color
				Prods  []product.Product
				Pgn    Pagination
				Qstr   Qstr
			}{
				Header: Header{
					Auth: r.Context().Value(helper.AuthUserKey).(user.User),
					Link: stype,
				},
				Ctgs:   ctgs,
				Brands: brands,
				Sizes:  sizes,
				Colors: colors,
				Prods:  prods,
				Pgn:    pgn,
				Qstr:   qstr,
			}

			helper.Render(w, "shop.gohtml", data)
			return
		}

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	})
}

func Details() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/product/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			id, _ := strconv.Atoi(r.FormValue("id"))
			prod, err := product.FindOne(id)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			images, err := file.FindByProductID(id)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			data := struct {
				Header Header
				Prod   product.Product
				Images []file.File
			}{
				Header: Header{
					Auth: r.Context().Value(helper.AuthUserKey).(user.User),
				},
				Prod:   prod,
				Images: images,
			}

			helper.Render(w, "product_details.gohtml", data)
			return
		}

		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	})
}

func createPgn(page int, quantity int) Pagination {
	pgn := Pagination{}
	pages := math.Ceil(float64(quantity) / 9)
	for i := 1; i <= int(pages); i++ {
		switch i {
		case page - 2:
			pgn.Bprev = i
		case page - 1:
			pgn.Prev = i
		case page:
			pgn.Current = i
		case page + 1:
			pgn.Next = i
		case page + 2:
			pgn.Anext = i
		}
	}

	return pgn
}

func arrangeQstr(r *http.Request) Qstr {
	qstr := Qstr{}

	if r.FormValue("g") != "" {
		qstr.Gender = "&g=" + r.FormValue("g")
	}
	ctgID, _ := strconv.Atoi(r.FormValue("ctg"))
	if ctgID > 0 {
		qstr.Ctg = fmt.Sprintf("&ctg=%d", ctgID)
	}
	if r.FormValue("b") != "" {
		qstr.Brands = "&b=" + r.FormValue("b")
	}
	if r.FormValue("s") != "" {
		qstr.Sizes = "&s=" + r.FormValue("s")
	}
	if r.FormValue("c") != "" {
		qstr.Colors = "&c=" + r.FormValue("c")
	}

	return qstr
}
