package main

import (
	"log"
	"os"

	"github.com/co0p/gopherconeu-is/pkg/routing"
	"github.com/co0p/gopherconeu-is/pkg/webserver"
	"github.com/co0p/gopherconeu-is/version"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	shutdown := make(chan error, 2)

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

	br := routing.BaseRouter()
	ws := webserver.New("", port, br)

	go func() {
		err := ws.Start()
		log.Fatal(err)
	}()

	dr := routing.DiagnosticsRouter()
	dws := webserver.New("", internalPort, dr)

	go func() {
		err := dws.Start()
		log.Fatal(err)
	}()

	select {
	case killSignal := <-interrupt:
		log.Printf("Got signal %s", killSignal)

	case err := <-shutdown:
		log.Printf("shutdown because %s...", err.Error())
		dws.Stop()
		ws.Stop()
	}
}
