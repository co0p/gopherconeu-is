package main

import (
	"log"
	"os"

	"github.com/co0p/gopherconeu-is/pkg/routing"
	"github.com/co0p/gopherconeu-is/pkg/webserver"
)

func main() {
	log.Printf("Service is starting ...")

	port := os.Getenv("PORT")
	log.Printf("using port %s", port)
	if len(port) < 1 {
		log.Fatal("PORT was not specified")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	log.Fatal(ws.Start())

}
