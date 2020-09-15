package main

import (
	"context"
	"log"
	"net/http"
	"onlineshop/admin/product"
	"onlineshop/app/blog"
	"onlineshop/app/cart"
	"onlineshop/app/contact"
	"onlineshop/app/home"
	"onlineshop/app/shop"
	"onlineshop/app/user"
	"onlineshop/helper"
)

func main() {
	http.Handle("/", basic(home.Index()))
	http.Handle("/shop", basic(shop.Index()))
	http.Handle("/cart", basic(cart.Index()))
	http.Handle("/checkout", basic(cart.Checkout()))
	http.Handle("/blog", basic(blog.Index()))
	http.Handle("/contact", basic(contact.Index()))
	http.Handle("/login", user.Login())
	http.Handle("/logout", user.Logout())
	http.Handle("/register", user.Register())
	http.Handle("/admin", basic(product.Index()))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./static"))))

	log.Println("Server running...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Basic middleware
func basic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _ := user.GetAuthUser(r)
		if u.ID > 0 {
			cookie, _ := r.Cookie("session_id")
			cookie.MaxAge = 15
			http.SetCookie(w, cookie)
		}

		ctx := context.WithValue(r.Context(), helper.AuthUserKey, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
