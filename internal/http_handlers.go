package internal

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func GenerateTestcaseSenariosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.PostFormValue("explanation"))
}

func StreamingGenerateTestcaseSenariosHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")

	stream, _ := StreamingGenerateTestcases("A REST API to fetch a list of users.")
	for chunk := range stream {
		fmt.Fprint(w, chunk)
		flusher.Flush()
	}
}
