package configs

import (
	"log/slog"
	"sync"

	env "github.com/caarlos0/env/v11"
)

type Config struct {
	Debug              bool   `env:"DEBUG" envDefault:"false"`
	LogLevel           string `env:"LOG_LEVEL" envDefault:"info"`
	PostgresDNS        string `env:"POSTGRES_DNS" envDefault:"postgres://calapi:calapi@localhost:5432/calapi_db?sslmode=disable"`
	HTTPListenHostPort string `env:"HTTP_LISTEN_HOST_PORT" envDefault:"0.0.0.0:2090"`
}

var instance Config
var instanceOnce sync.Once

// Get returns the config instance
func Get() *Config {
	instanceOnce.Do(func() {
		if err := env.Parse(&instance); err != nil {
			slog.Error("unable to load secrets from .env", "error", err)
			panic(err)
		}
	})
	return &instance
}

// // this will only contain keys that doesn't have a default value
// // and will be used only for tests to load those values

// func LoadConfigForTest() {
// 	instanceOnce.Do(func() {
// 		instance = Config{
// 			Debug:               true,
// 			LogLevel:            "debug",
// 			PostgresDNS:         "postgres://calapi:calapi@localhost:5432/calapi_db?sslmode=disable",
// 			HTTPListenHostPort:  "0.0.0.0:2090",
// 		}
// 		slog.Info("printing the configs", "config", instance)
// 	})
// }
