package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHTTPPort      = "8080"
	defaultHTTPRWTimeout = 10 * time.Second
)

type (
	Config struct {
		ServerConfig HttpConfig
	}

	HttpConfig struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
)

func NewConfig(path string) (*Config, error) {
	setDefaults()
	cfg := &Config{}
	if err := parseConfig(path); err != nil {
		return nil, err
	}

	if err := unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func setDefaults() {
	viper.SetDefault("interface.port", defaultHTTPPort)
	viper.SetDefault("interface.readTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("interface.writeTimeout", defaultHTTPRWTimeout)
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("server", &cfg.ServerConfig); err != nil {
		return err
	}
	return nil
}
