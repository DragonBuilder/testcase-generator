package main

import (
	"fmt"
	"log"
	"net/http"
	"testcase-generator/internal"
	"testcase-generator/internal/config"
)

func main() {
	config.Init()
	// http.Handle("/static", http.FileServer(http.Dir("./static/")))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", internal.IndexHandler)
	http.HandleFunc("/generate/scenarios", internal.GenerateTestcaseSenariosHandler)
	// http.HandleFunc("/generate/scenarios/streaming", internal.StreamingGenerateTestcaseSenariosHandler)
	// http.HandleFunc("/start/generate/scenarios/streaming", internal.StartStreamingGenerateTestcaseSenariosHandler)
	http.HandleFunc("/generate/scenarios/streaming", internal.StartStreamingGenerateTestcaseSenariosHandler)

	log.Printf("Starting app on port: %s\n", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), nil))
}
