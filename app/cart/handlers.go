package cart

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"onlineshop/app/user"
	"onlineshop/helper"
)

// Header struct
type Header struct {
	Auth user.User
	Link string
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
					Auth: r.Context().Value(helper.AuthUserKey).(user.User),
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

			exists, err := uc.store()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			msg := "The item successfully added to your cart"
			if exists {
				msg = "The item is already in the cart"
			}

			j, err := json.Marshal(msg)
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
					Auth: r.Context().Value(helper.AuthUserKey).(user.User),
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
