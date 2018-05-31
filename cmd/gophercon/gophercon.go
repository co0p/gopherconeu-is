package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "hi there")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server %v", err.Error())
	}
}
