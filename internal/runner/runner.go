package runner

import (
	"calendar/internal/config"
	"calendar/internal/server"
	log "github.com/sirupsen/logrus"
)

func Start(path string) {
	cfg := creteConfig(path)
	startServer(cfg)
}

func creteConfig(path string) *config.Config {
	log.Info("Start parse config")
	cfg, err := config.NewConfig(path)
	if err != nil {
		log.WithError(err).Fatal("failed to parse config")
	}
	return cfg
}

func startServer(cfg *config.Config) {
	log.Infof("Starting server in port: %s", cfg.ServerConfig.Port)
	s := server.New(cfg.ServerConfig, nil)

	if err := s.Run(); err != nil {
		log.WithError(err).Fatal("failed start sever")
	}
}
