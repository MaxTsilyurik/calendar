package server

import (
	"calendar/internal/config"
	"context"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New(config config.HttpConfig, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + config.Port,
			Handler:      handler,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
