package home

import (
	"io"
	"net/http"
	"onlineshop/admin/product"
	"onlineshop/helper"
)

// Header struct
type Header struct {
	Context helper.ContextData
	Link    string
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			prods, err := product.FindNew()
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}

			data := struct {
				Header Header
				Prods  []product.Product
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
					Link:    "home",
				},
				Prods: prods,
			}

			helper.Render(w, "home.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
