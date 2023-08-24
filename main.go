package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	log.Printf("Starting app on port: %s\n", Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Config.Port), nil))
}
