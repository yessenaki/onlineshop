package config

import (
	"html/template"
	"log"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	Tpl = template.Must(Tpl.ParseGlob("templates/admin/*.gohtml"))
	log.Println("Tpl connected")
}
