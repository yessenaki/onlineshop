package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/cart", cart)
	http.HandleFunc("/checkout", checkout)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "GET /")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func shop(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "GET /shop")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /shop")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func cart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "GET /cart")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /cart")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "GET /checkout")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /checkout")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func blog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "GET /blog")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /blog")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "GET /contact")
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /contact")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}
