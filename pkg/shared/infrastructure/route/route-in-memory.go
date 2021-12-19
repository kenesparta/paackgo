package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitInMemoryRoutes(r *mux.Router) {
	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
		}).Methods("GET")
}
