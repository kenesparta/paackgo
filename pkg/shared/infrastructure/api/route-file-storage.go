package api

import (
	"github.com/gorilla/mux"
	apiCity "github.com/kenesparta/paackgo/entities/city/infrastructure/api"
	"github.com/kenesparta/paackgo/shared/infrastructure/persistence"
)

func InitFileStorageRoutes(s *persistence.FileRepository, r *mux.Router) {
	apiCity.NewCityHandler(s, r)
}
