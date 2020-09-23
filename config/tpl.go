package config

import (
	"fmt"
	"html/template"
	"log"
)

var Tpl *template.Template

func init() {
	funcMap := template.FuncMap{
		"fmtprice": priceToString,
	}

	Tpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.gohtml"))
	Tpl = template.Must(Tpl.ParseGlob("templates/admin/*.gohtml"))
	log.Println("Tpl connected")
}

func priceToString(price int) string {
	fprice := (float64(price) + 0.1) / 100
	return fmt.Sprintf("%.2f", fprice)
}
