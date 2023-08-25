package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	data, err := io.ReadAll(r.Body)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error reading request body: %v", err))
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		os.Stderr.WriteString("No data provided")
		http.Error(w, "No question in body!", http.StatusInternalServerError)
		return
	}
	// data := r.PostFormValue("explanation")
	log.Println(string(data))

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")

	// stream, _ := StreamingGenerateTestcases("A REST API to fetch a list of users.")
	stream, _ := StreamingGenerateTestcases(string(data))
	// stream, _ := StreamingGenerateTestcases(string(data))
	for chunk := range stream {
		content := fmt.Sprintf(`data: {"content":"%s"}\n\n`, chunk.Choices[0].Delta.Content)

		// log.Println(content)
		fmt.Fprint(w, content)
		// fmt.Fprint(w, chunk.Choices[0].Delta.Content)
		flusher.Flush()
	}
}
