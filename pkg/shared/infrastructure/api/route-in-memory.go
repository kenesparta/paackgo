package api

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/shared/infrastructure/persistence"
	"net/http"
)

func InitInMemoryRoutes(s *persistence.MemoryRepository, r *mux.Router) {
	r.HandleFunc("/v1/", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
}
