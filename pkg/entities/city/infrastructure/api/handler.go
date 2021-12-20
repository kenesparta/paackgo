package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kenesparta/paackgo/entities/city/application"
	"github.com/kenesparta/paackgo/shared/infrastructure/header"
	"log"
	"net/http"
)

type City struct {
	app application.CityAppInterface
}

func NewCity(app application.CityAppInterface) *City {
	return &City{app: app}
}

func (c *City) Get(w http.ResponseWriter, r *http.Request) {
	header.Headers(w)
	cityId, err := application.TransformToCityId(mux.Vars(r)["id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	exReq, err := c.app.Find(r.Context(), *cityId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&exReq); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
