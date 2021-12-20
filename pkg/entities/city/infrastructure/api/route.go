package api

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/shared/infrastructure/persistence"
	"net/http"
)

func NewCityHandler(fr *persistence.FileRepository, r *mux.Router) {
	c := NewCity(fr.City)
	r.HandleFunc("/v1/city/{id}", c.Get).Methods(http.MethodGet)
}
