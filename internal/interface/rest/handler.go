package rest

import (
	"calendar/internal/app"
	v1 "calendar/internal/interface/rest/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewHandler(app app.Application) http.Handler {
	router := chi.NewRouter()
	addMiddlewares(router)
	apiRoute := chi.NewRouter()
	apiRoute.Mount("/api/v1", v1.NewHandler(app, router))
	return apiRoute
}

func addMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)
	router.Use(NewStructuredLogger(logrus.StandardLogger()))
}
