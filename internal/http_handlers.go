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
	tmpl := template.Must(template.ParseFiles("templates/_streaming_chat.html"))
	tmpl.Execute(w, nil)
}

// const x = `
// <div hx-ext="sse" sse-swap="message" hx-swap="beforeend">
//     %s
// </div>
// `

func StartStreamingGenerateTestcaseSenariosHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")

	// stream, _ := StreamingGenerateTestcases("A REST API to fetch a list of users.")
	stream, _ := StreamingGenerateTestcases("What are you?")
	for chunk := range stream {
		// fmt.Print(chunk.Choices[0].Delta.Content)
		content := fmt.Sprintf("data: [%s]\n\n", chunk.Choices[0].Delta.Content)

		// log.Println(content)
		fmt.Fprint(w, content)
		flusher.Flush()
	}
}
