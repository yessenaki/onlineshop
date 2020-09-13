package config

import (
	"html/template"
	"log"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	log.Println("Tpl connected")
}
