package configuration

import (
	"time"

	xlogger "github.com/syth0le/gopnik/logger"
	xservers "github.com/syth0le/gopnik/servers"
)

const (
	defaultAppName = "starter"
)

func NewDefaultConfig() *Config {
	return &Config{
		Logger: xlogger.LoggerConfig{
			Level:       xlogger.InfoLevel,
			Encoding:    "console",
			Path:        "stdout",
			Environment: xlogger.Development,
		},
		Application: ApplicationConfig{
			GracefulShutdownTimeout: 15 * time.Second,
			ForceShutdownTimeout:    20 * time.Second,
			App:                     defaultAppName,
		},
		PublicServer: xservers.ServerConfig{
			Enable:   false,
			Endpoint: "",
			Port:     0,
		},
		Temporal: TemporalConfig{
			Enable:   false,
			Endpoint: "",
		},
	}
}
