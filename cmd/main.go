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
	http.HandleFunc("/", internal.Index)
	http.HandleFunc("/testcase/generate", internal.Generate)

	log.Printf("Starting app on port: %s\n", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), nil))
}
