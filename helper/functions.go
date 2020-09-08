package helper

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

func RenderTemplate(w http.ResponseWriter, path map[string]string, ctx interface{}) {
	t := template.Must(template.ParseGlob("templates/layouts/*.gohtml"))
	t = template.Must(t.ParseGlob("templates/" + path["folder"] + "/*.gohtml"))
	err := t.ExecuteTemplate(w, path["file"], ctx)
	throwError(w, err)
}

func throwError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func currentTime() string {
	t := time.Now()
	datetime := t.Format("2006-01-02 15:04:05")
	return datetime
}
