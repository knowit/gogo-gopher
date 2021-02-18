package config

import "github.com/knowit/gogo-gopher/crud-webapp/pkg/runenv"


func New() *Config {
	return &Config{
		Environment: EnvironmentConfig{
			production: runenv.GetEnvAsBool("IS_PRODUCTION", false),
		},
		DebugMode: runenv.GetEnvAsBool("DEBUG_MODE", true),
	}
}

