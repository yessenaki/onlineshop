package shop

import (
	"io"
	"net/http"
	"onlineshop/admin/category"
	"onlineshop/app/user"
	"onlineshop/helper"
)

// Relative = relative categories
type Relative struct {
	Parent category.Category
	Childs []category.Category
}

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/shop/" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		auth := helper.AuthUserFromContext(r.Context())

		if r.Method == http.MethodGet {
			gender := 2
			isKids := 2
			if r.FormValue("t") == "men" {
				gender = 1
			} else if r.FormValue("t") == "kids" {
				isKids = 1
			}

			ctgs, err := category.ByGenderAndKids(gender, isKids)
			if err != nil {
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			}

			data := struct {
				Auth user.User
				Rlts []Relative
			}{
				Auth: auth,
				Rlts: createRelation(ctgs),
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

func createRelation(ctgs []category.Category) []Relative {
	rlts := []Relative{}
	for _, ctg := range ctgs {
		if ctg.ParentID == 0 {
			rlt := Relative{}
			rlt.Parent = ctg
			childs := []category.Category{}

			for _, ctg2 := range ctgs {
				if ctg2.ParentID == ctg.ID {
					childs = append(childs, ctg2)
				}
			}

			rlt.Childs = childs
			rlts = append(rlts, rlt)
		}
	}

	return rlts
}
