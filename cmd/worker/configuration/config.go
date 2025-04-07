package configuration

import (
	"time"

	"temporal-docs/internal/model"

	xlogger "github.com/syth0le/gopnik/logger"
)

type Config struct {
	Logger      xlogger.LoggerConfig `yaml:"logger"`
	Application ApplicationConfig    `yaml:"application"`
	Temporal    TemporalConfig       `yaml:"temporal"`
}

type ApplicationConfig struct {
	GracefulShutdownTimeout time.Duration      `yaml:"graceful_shutdown_timeout"`
	ForceShutdownTimeout    time.Duration      `yaml:"force_shutdown_timeout"`
	App                     string             `yaml:"app"`
	Type                    model.ScheduleType `yaml:"type"`
}

type TemporalConfig struct {
	Enable   bool   `yaml:"enable"`
	Endpoint string `yaml:"endpoint"`
}
