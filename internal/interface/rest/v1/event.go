package v1

import (
	"calendar/internal/app"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	body, ok := readBody(w, r)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var e app.CreateEventCommand
	if err := jsoniter.Unmarshal(body, &e); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eventId, err := h.app.Commands.CreateEvent.Handle(r.Context(), e)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Location", fmt.Sprintf("/events/%s", eventId))
	w.WriteHeader(http.StatusCreated)
}

func (h handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {

}
func (h handler) GetEventByDay(w http.ResponseWriter, r *http.Request) {

}

func (h handler) GetEventByWeek(w http.ResponseWriter, r *http.Request) {

}

func (h handler) GetEventByMonth(w http.ResponseWriter, r *http.Request) {

}
