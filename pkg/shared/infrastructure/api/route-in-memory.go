package api

import (
	"github.com/gorilla/mux"
	apiTrip "github.com/kenesparta/paackgo/entities/trip/infrastructure/api"
	"github.com/kenesparta/paackgo/logger"
	"github.com/kenesparta/paackgo/shared/infrastructure/persistence"
	"net/http"
)

func InitInMemoryRoutes(s *persistence.MemoryRepository, fs *persistence.FileRepository, r *mux.Router, l *logger.NewLogger) {
	r.HandleFunc("/v1/", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	apiTrip.NewCityHandler(s, fs, r, l)
}
