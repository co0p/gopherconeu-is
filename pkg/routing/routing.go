package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BaseRouter configures the /home handler and returns the configured router
func BaseRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)
	return r
}

func homeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("processing request ...%s", r.URL.Path)
		fmt.Fprint(w, "hi there")
	}
}
