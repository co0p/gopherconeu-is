package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Printf("Service is starting ...")
	http.HandleFunc("/home", homeHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server %v", err.Error())
	}
}

func homeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("processing request ...%s", r.URL.Path)
		fmt.Fprint(w, "hi there")
	}
}
