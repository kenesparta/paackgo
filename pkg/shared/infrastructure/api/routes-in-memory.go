package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitInMemoryRoutes(r *mux.Router) {
	r.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
		}).Methods("GET")
}
