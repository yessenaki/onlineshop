package helper

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, path map[string]string, ctx interface{}) {
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
