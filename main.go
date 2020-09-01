package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/cart", cart)
	http.HandleFunc("/checkout", checkout)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/contact", contact)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Home Page"
		renderTemplate(w, "home", ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func shop(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Shop Page"
		renderTemplate(w, "shop", ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /shop")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func cart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Cart Page"
		renderTemplate(w, "cart", ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /cart")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Checkout Page"
		renderTemplate(w, "checkout", ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /checkout")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func blog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Blog Page"
		renderTemplate(w, "blog", ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /blog")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Contact Page"
		renderTemplate(w, "contact", ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /contact")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func renderTemplate(w http.ResponseWriter, folder string, ctx interface{}) {
	t := template.Must(template.ParseGlob("templates/layouts/*.gohtml"))
	t = template.Must(t.ParseGlob("templates/" + folder + "/*.gohtml"))
	err := t.ExecuteTemplate(w, "index.gohtml", ctx)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
