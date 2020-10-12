package contact

import (
	"io"
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
		if r.URL.Path != "/contact/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		if r.Method == http.MethodGet {
			data := struct {
				Header Header
			}{
				Header: Header{
					Context: helper.GetContextData(r.Context()),
					Link:    "contact",
				},
			}

			helper.Render(w, "contact.gohtml", data)
			return
		} else if r.Method == http.MethodPost {
			io.WriteString(w, "POST /")
		} else {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	})
}
