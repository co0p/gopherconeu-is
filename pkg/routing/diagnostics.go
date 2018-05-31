package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DiagnosticsRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthzHandler()).Methods(http.MethodGet)
	return r
}

func healthzHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
