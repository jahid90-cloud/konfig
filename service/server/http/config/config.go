package config

import (
	"github.com/jahid90-cloud/konfig/service/server/http/apps"
	"github.com/jahid90-cloud/konfig/service/server/http/apps/configuration"
	"github.com/jahid90-cloud/konfig/service/server/http/apps/ping"
	"github.com/jahid90-cloud/konfig/service/utils/env"
)

type Config struct {
	Env  env.Env
	Apps []apps.App
}

func NewConfig(env env.Env) *Config {
	pingApp := ping.App()
	configurationApp := configuration.App()

	apps := []apps.App{pingApp, configurationApp}

	return &Config{
		Apps: apps,
		Env:  env,
	}
}
