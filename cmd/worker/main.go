package main

import (
	"fmt"
	"log"

	"temporal-docs/cmd/worker/application"
	"temporal-docs/cmd/worker/configuration"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/pflag"
	xlogger "github.com/syth0le/gopnik/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	// TODO: validate

	logger, err := constructLogger(cfg.Logger)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	logger.Sugar().Infof("worker config %#v", cfg)

	app := application.New(cfg, logger)
	if err = app.Run(); err != nil {
		logger.Sugar().Fatalf("application stopped with error: %v", err)
	} else {
		logger.Info("application stopped")
	}
}

func constructLogger(cfg xlogger.LoggerConfig) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	switch cfg.Environment {
	case xlogger.Development:
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, fmt.Errorf("new development logger")
		}
	case xlogger.Production:
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("new production logger")
		}
	default:
		return nil, fmt.Errorf("unexpected environment for logger: %w", err)
	}
	logger = logger.With(zap.String("env", string(cfg.Environment)))

	defer logger.Sync()
	return logger, nil
}

func loadConfig() (*configuration.Config, error) {
	cfg := configuration.NewDefaultConfig()

	configPath := pflag.StringP("config", "c", "", "config path")
	pflag.Parse()

	if err := cleanenv.ReadConfig(*configPath, cfg); err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}
	return cfg, nil
}
