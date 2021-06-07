package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

}

func MakeHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	return r
}
