package main

import (
	"log"
	"net/http"

	"github.com/co0p/gopherconeu-is/pkg/routing"
)

func main() {
	log.Printf("Service is starting ...")

	r := routing.BaseRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server %v", err.Error())
	}
}
