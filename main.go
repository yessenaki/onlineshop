package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"onlineshop/app/user"
	"onlineshop/config"

	_ "github.com/lib/pq"
)

func main() {
	config.InitDB("postgres://postgres:postgres@localhost/onlineshop?sslmode=disable")

	http.HandleFunc("/", home)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/cart", cart)
	http.HandleFunc("/checkout", checkout)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/logout", user.Logout)
	http.HandleFunc("/register", user.Register)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./static"))))

	log.Println("Server running...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Home Page"
		var path = map[string]string{
			"folder": "home",
			"file":   "index.gohtml",
		}
		renderTemplate(w, path, ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func shop(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Shop Page"
		var path = map[string]string{
			"folder": "shop",
			"file":   "index.gohtml",
		}
		renderTemplate(w, path, ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /shop")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func cart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Cart Page"
		var path = map[string]string{
			"folder": "cart",
			"file":   "index.gohtml",
		}
		renderTemplate(w, path, ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /cart")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Checkout Page"
		var path = map[string]string{
			"folder": "checkout",
			"file":   "index.gohtml",
		}
		renderTemplate(w, path, ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /checkout")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func blog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Blog Page"
		var path = map[string]string{
			"folder": "blog",
			"file":   "index.gohtml",
		}
		renderTemplate(w, path, ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /blog")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ctx interface{} = "Contact Page"
		var path = map[string]string{
			"folder": "contact",
			"file":   "index.gohtml",
		}
		renderTemplate(w, path, ctx)
	} else if r.Method == http.MethodPost {
		io.WriteString(w, "POST /contact")
	} else {
		http.Error(w, "405 method not allowed", 405)
	}
}

func renderTemplate(w http.ResponseWriter, path map[string]string, ctx interface{}) {
	t := template.Must(template.ParseGlob("templates/layouts/*.gohtml"))
	t = template.Must(t.ParseGlob("templates/" + path["folder"] + "/*.gohtml"))
	err := t.ExecuteTemplate(w, path["file"], ctx)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
