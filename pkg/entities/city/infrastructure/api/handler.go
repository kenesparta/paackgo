package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/entities/city/application"
	"github.com/kenesparta/paackgo/entities/city/domain"
	"github.com/kenesparta/paackgo/logger"
	"github.com/kenesparta/paackgo/shared/infrastructure/header"
	"net/http"
)

type City struct {
	app    application.CityAppInterface
	logger *logger.NewLogger
}

func NewCity(app application.CityAppInterface, logger *logger.NewLogger) *City {
	return &City{app: app, logger: logger}
}

func (c *City) Get(w http.ResponseWriter, r *http.Request) {
	header.Headers(w)
	cityId, err := domain.TransformToCityId(mux.Vars(r)["id"])
	if err != nil {
		c.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	cityReq, err := c.app.Find(r.Context(), *cityId)
	if err != nil {
		c.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&cityReq); err != nil {
		c.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
