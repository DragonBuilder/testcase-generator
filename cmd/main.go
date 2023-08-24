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
	http.HandleFunc("/", internal.IndexHandler)
	http.HandleFunc("/generate/scenarios", internal.GenerateTestcaseSenariosHandler)
	http.HandleFunc("/generate/scenarios/streaming", internal.StreamingGenerateTestcaseSenariosHandler)

	log.Printf("Starting app on port: %s\n", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), nil))
}
