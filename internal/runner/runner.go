package runner

import (
	"calendar/internal/adapter/repository/mock"
	"calendar/internal/app"
	"calendar/internal/app/command"
	"calendar/internal/app/query"
	"calendar/internal/config"
	"calendar/internal/interface/rest"
	"calendar/internal/server"
	log "github.com/sirupsen/logrus"
)

func Start(path string) {
	cfg := creteConfig(path)
	application := newApplication()
	startServer(cfg, application)
}

func newApplication() app.Application {
	repo := mock.NewInMemoryEventRepository()

	return app.Application{
		Commands: app.Commands{
			CreateEvent: command.NewCreateEventHandler(repo),
			EditEvent:   nil,
			RemoveEvent: command.NewDeleteEventHandler(repo),
		},
		Queries: app.Queries{
			GetEventForDay:   query.NewEventForDayHandler(repo),
			GetEventOfWeek:   query.NewEventForWeekHandler(repo),
			GetEventForMonth: query.NewEventForMonthHandler(repo),
		},
	}
}

func creteConfig(path string) *config.Config {
	log.Info("Start parse config")
	cfg, err := config.NewConfig(path)
	if err != nil {
		log.WithError(err).Fatal("failed to parse config")
	}
	return cfg
}

func startServer(cfg *config.Config, app app.Application) {
	log.Infof("Starting server in port: %s", cfg.ServerConfig.Port)
	s := server.New(cfg.ServerConfig, rest.NewHandler(app))
	if err := s.Run(); err != nil {
		log.WithError(err).Fatal("failed start sever")
	}
}
