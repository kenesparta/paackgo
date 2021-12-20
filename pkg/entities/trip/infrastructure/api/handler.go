package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	appCity "github.com/kenesparta/paackgo/entities/city/application"
	appTrip "github.com/kenesparta/paackgo/entities/trip/application"
	"github.com/kenesparta/paackgo/entities/trip/domain"
	"github.com/kenesparta/paackgo/logger"
	"github.com/kenesparta/paackgo/shared/infrastructure/header"
	"net/http"
)

type Trip struct {
	appTrip appTrip.TripAppInterface
	appCity appCity.CityAppInterface
	logger  *logger.NewLogger
}

func NewTrip(appTrip appTrip.TripAppInterface, appCity appCity.CityAppInterface, logger *logger.NewLogger) *Trip {
	return &Trip{appTrip, appCity, logger}
}

func (t *Trip) Save(w http.ResponseWriter, r *http.Request) {
	header.Headers(w)
	var tr TripRequest
	if err := json.NewDecoder(r.Body).Decode(&tr); nil != err {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := tr.Validate(r.Context(), t.appCity); err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tripId, err := t.appTrip.Save(r.Context(), *tr.ParseToTrip())
	if err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"tripId":"%d"}`, *tripId)))
}

func (t *Trip) Get(w http.ResponseWriter, r *http.Request) {
	header.Headers(w)
	tripResponse := TripCity{
		ctx: r.Context(),
		app: t.appCity,
	}
	tripId, err := domain.TransformToTripId(mux.Vars(r)["id"])
	if err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	tripReq, err := t.appTrip.Find(r.Context(), *tripId)
	if err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	tripResp, err := tripResponse.ParseToTrip(tripReq)
	if err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&tripResp); err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *Trip) GetAll(w http.ResponseWriter, r *http.Request) {
	header.Headers(w)
	tripResponse := TripCity{
		ctx: r.Context(),
		app: t.appCity,
	}
	trips, err := t.appTrip.FindAll(r.Context())
	if err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	tripResp, err := tripResponse.ParseToListOfTrips(trips)
	if err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&tripResp); err != nil {
		t.logger.ErrorLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
