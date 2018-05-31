package main

import (
	"log"
	"os"

	"github.com/co0p/gopherconeu-is/pkg/routing"
	"github.com/co0p/gopherconeu-is/pkg/webserver"
	"github.com/co0p/gopherconeu-is/version"
)

func main() {
	log.Printf("Service %s (Version:%s, BuildTime:%s) is starting ...", version.Release, version.Commit, version.BuildTime)

	port := os.Getenv("PORT")
	log.Printf("using port %s", port)
	if len(port) < 1 {
		log.Fatal("PORT was not specified")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	log.Fatal(ws.Start())

}
