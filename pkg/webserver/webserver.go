package webserver

import (
	"net"
	"net/http"
)

type WebServer struct {
	http.Server
}

func New(host, port string, h http.Handler) *WebServer {
	var ws WebServer
	ws.Addr = net.JoinHostPort(host, port)
	ws.Handler = h

	return &ws
}

func (ws *WebServer) Start() error {
	return ws.ListenAndServe()
}
