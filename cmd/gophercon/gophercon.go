package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Service is starting ...")

	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server %v", err.Error())
	}
}

func homeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("processing request ...%s", r.URL.Path)
		fmt.Fprint(w, "hi there")
	}
}
