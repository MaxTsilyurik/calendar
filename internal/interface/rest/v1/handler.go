package v1

import (
	"calendar/internal/app"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ServerInterface interface {
	CreateEvent(writer http.ResponseWriter, request *http.Request)
	DeleteEvent(writer http.ResponseWriter, request *http.Request)
	GetEventByDay(writer http.ResponseWriter, request *http.Request)
	GetEventByMonth(writer http.ResponseWriter, request *http.Request)
	GetEventByWeek(writer http.ResponseWriter, request *http.Request)
}

type handler struct {
	app app.Application
}

func NewHandler(app app.Application, r *chi.Mux) http.Handler {
	return NewMuxServer(
		handler{app: app},
		r,
	)
}

func NewMuxServer(s ServerInterface, r *chi.Mux) http.Handler {
	if r == nil {
		r = chi.NewRouter()
	}

	r.Group(func(r chi.Router) {
		r.Post("/events", s.CreateEvent)
		r.Delete("/events/{id}", s.DeleteEvent)
		r.Get("/events/by-day", s.GetEventByDay)
		r.Get("/events/by-week", s.GetEventByWeek)
		r.Get("/events/by-month", s.GetEventByMonth)
	})

	return r
}
