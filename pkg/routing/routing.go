package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

// BaseRouter configures the /home handler and returns the instance
func BaseRouter() *mux.Router {
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)
	return r
}

func homeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("processing request ...%s", r.URL.Path)
		fmt.Fprint(w, "hi there")
	}
}
