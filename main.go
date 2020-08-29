package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/cart", cart)
	http.HandleFunc("/checkout", checkout)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/contact", contact)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./resources"))))

	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "index.gohtml", nil)
		handleError(w, err)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func shop(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "shop.gohtml", nil)
		handleError(w, err)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /shop")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func cart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "cart.gohtml", nil)
		handleError(w, err)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /cart")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "checkout.gohtml", nil)
		handleError(w, err)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /checkout")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func blog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "blog.gohtml", nil)
		handleError(w, err)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /blog")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "contact.gohtml", nil)
		handleError(w, err)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /contact")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
