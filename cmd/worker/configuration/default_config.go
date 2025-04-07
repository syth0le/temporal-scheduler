package configuration

import (
	"time"

	"temporal-docs/internal/model"

	xlogger "github.com/syth0le/gopnik/logger"
)

const (
	defaultAppName = "starter"
	defaultType    = model.ColdScheduleType
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
			Type:                    defaultType,
		},
		Temporal: TemporalConfig{
			Enable:   false,
			Endpoint: "",
		},
	}
}
