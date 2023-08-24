package main

import (
	"fmt"
	"log"
	"net/http"
	"testcase-generator/internal"
)

func main() {
	http.HandleFunc("/", internal.Index)

	http.HandleFunc("/testcase/generate", internal.Generate)

	log.Printf("Starting app on port: %s\n", Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Config.Port), nil))
}
