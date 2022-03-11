package config

import (
	"github.com/jahid90-cloud/konfig/service/server/http/apps"
	"github.com/jahid90-cloud/konfig/service/server/http/apps/ping"
)

type Config struct {
	Apps []apps.App
}

func NewConfig() *Config {
	pingApp := ping.App()

	apps := []apps.App{pingApp}

	return &Config{
		Apps: apps,
	}
}
