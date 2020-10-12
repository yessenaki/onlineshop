package cart

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"onlineshop/helper"
)

// Header struct
type Header struct {
	Context helper.ContextData
	Link    string
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/cart/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			data := struct {
				Header Header
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
				},
			}

			helper.Render(w, "cart.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			var uc UserCart
			err := helper.DecodeJSONBody(w, r, &uc)
			if err != nil {
				var mr *helper.MalformedRequest
				if errors.As(err, &mr) {
					http.Error(w, mr.Msg, mr.Status)
				} else {
					log.Println(err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			data := struct {
				Status  bool
				Message string
			}{
				Status:  true,
				Message: "The item successfully added to your cart",
			}

			exists, err := uc.store()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if exists {
				data.Status = false
				data.Message = "The item is already in the cart"
			}

			j, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(j)
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}

func Checkout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := struct {
				Header Header
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
				},
			}

			helper.Render(w, "checkout.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
