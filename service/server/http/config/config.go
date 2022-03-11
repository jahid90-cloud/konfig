package config

import (
	"github.com/jahid90-cloud/konfig/service/server/http/apps"
	"github.com/jahid90-cloud/konfig/service/server/http/apps/configuration"
	"github.com/jahid90-cloud/konfig/service/server/http/apps/ping"
)

type Config struct {
	Apps []apps.App
}

func NewConfig() *Config {
	pingApp := ping.App()
	configurationApp := configuration.App()

	apps := []apps.App{pingApp, configurationApp}

	return &Config{
		Apps: apps,
	}
}
