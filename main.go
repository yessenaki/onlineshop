package main

import (
	"context"
	"log"
	"net/http"
	"onlineshop/admin/brand"
	"onlineshop/admin/category"
	"onlineshop/admin/color"
	"onlineshop/admin/post"
	pc "onlineshop/admin/post/category"
	"onlineshop/admin/post/tag"
	"onlineshop/admin/product"
	"onlineshop/admin/size"
	"onlineshop/app/blog"
	"onlineshop/app/cart"
	"onlineshop/app/contact"
	"onlineshop/app/home"
	"onlineshop/app/shop"
	"onlineshop/app/user"
	"onlineshop/helper"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", home.Index())
	mux.Handle("/shop/", shop.Index())
	mux.Handle("/product/", shop.Details())
	mux.Handle("/cart/", cart.Index())
	mux.Handle("/checkout/", cart.Checkout())
	mux.Handle("/blog/", blog.Index())
	mux.Handle("/contact/", contact.Index())
	mux.Handle("/login/", user.Login())
	mux.Handle("/logout/", user.Logout())
	mux.Handle("/register/", user.Register())
	mux.Handle("/admin/products/", override(product.Handle()))
	mux.Handle("/admin/products/delete-image/", product.DeleteImage())
	mux.Handle("/admin/categories/", override(category.Handle()))
	mux.Handle("/admin/brands/", override(brand.Handle()))
	mux.Handle("/admin/sizes/", override(size.Handle()))
	mux.Handle("/admin/colors/", override(color.Handle()))
	mux.Handle("/admin/post-categories/", override(pc.Handle()))
	mux.Handle("/admin/post-tags/", override(tag.Handle()))
	mux.Handle("/admin/posts/", override(post.Handle()))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static"))))

	log.Println("Server running...")
	err := http.ListenAndServe(":8080", basic(mux))
	if err != nil {
		log.Fatal(err)
	}
}

// Basic middleware
func basic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cd := helper.ContextData{}
		u, _ := user.GetAuthUser(r)
		if u.ID > 0 {
			cookie, _ := r.Cookie("session_id")
			cookie.Path = "/"
			cookie.MaxAge = 3600
			cookie.HttpOnly = true
			http.SetCookie(w, cookie)

			qnt, _ := cart.GetItemQuantity(u.ID)
			cd.Auth = u
			cd.ItemQnt = qnt
		}

		ctx := context.WithValue(r.Context(), helper.ContextDataKey, cd)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Method Override middleware
func override(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Look in the request body for a spoofed method.
			method := r.PostFormValue("_method")

			// Check that the spoofed method is a valid HTTP method and
			// update the request object accordingly.
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				r.Method = method
			}
		}

		next.ServeHTTP(w, r)
	})
}
