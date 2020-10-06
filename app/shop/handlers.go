package shop

import (
	"io"
	"net/http"
	"onlineshop/admin/brand"
	"onlineshop/admin/category"
	"onlineshop/admin/color"
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

			params := map[string]interface{}{
				"gender": gender,
				"isKids": isKids,
				"ctgID":  ctgID,
				"brands": r.FormValue("b"),
				"sizes":  r.FormValue("s"),
				"colors": r.FormValue("c"),
			}
			prods, err := product.ByParams(params)
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

			data := struct {
				Header Header
				Ctgs   []category.Parent
				Brands []brand.Brand
				Sizes  []size.Size
				Colors []color.Color
				Stype  string
				Prods  []product.Product
			}{
				Header: Header{
					Auth: helper.AuthUserFromContext(r.Context()),
					Link: stype,
				},
				Ctgs:   ctgs,
				Brands: brands,
				Sizes:  sizes,
				Colors: colors,
				Stype:  stype,
				Prods:  prods,
			}

			helper.Render(w, "shop.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
