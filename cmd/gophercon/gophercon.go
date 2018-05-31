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

	internalPort := os.Getenv("INTERNAL_PORT")
	log.Printf("using internal port %s", internalPort)
	if len(internalPort) < 1 {
		log.Fatal("INTERNAL_PORT was not specified")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	go func() {
		log.Fatal(ws.Start())
	}()

	dr := routing.DiagnosticsRouter()
	dws := webserver.New("", internalPort, dr)
	log.Fatal(dws.Start())
}
