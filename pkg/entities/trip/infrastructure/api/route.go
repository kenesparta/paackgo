package api

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/logger"
	"github.com/kenesparta/paackgo/shared/infrastructure/persistence"
	"net/http"
)

func NewCityHandler(mr *persistence.MemoryRepository, fr *persistence.FileRepository, r *mux.Router, l *logger.NewLogger) {
	t := NewTrip(mr.Trip, fr.City, l)
	r.HandleFunc("/v1/trip/{id}", t.Get).Methods(http.MethodGet)
	r.HandleFunc("/v1/trip", t.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/v1/trip", t.Save).Methods(http.MethodPost)
}
