package configuration

import (
	"time"

	xlogger "github.com/syth0le/gopnik/logger"
	xservers "github.com/syth0le/gopnik/servers"
)

type Config struct {
	Logger       xlogger.LoggerConfig  `yaml:"logger"`
	Application  ApplicationConfig     `yaml:"application"`
	PublicServer xservers.ServerConfig `yaml:"public_server"`
	Temporal     TemporalConfig        `yaml:"temporal"`
}

type ApplicationConfig struct {
	GracefulShutdownTimeout time.Duration `yaml:"graceful_shutdown_timeout"`
	ForceShutdownTimeout    time.Duration `yaml:"force_shutdown_timeout"`
	App                     string        `yaml:"app"`
}

type TemporalConfig struct {
	Enable   bool   `yaml:"enable"`
	Endpoint string `yaml:"endpoint"`
}
