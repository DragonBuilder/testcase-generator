package internal

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func Generate(w http.ResponseWriter, r *http.Request) {
	log.Println(r.PostFormValue("explanation"))
}
