package application

import (
	"context"
	"fmt"
	"syscall"

	"temporal-docs/cmd/starter/configuration"
	"temporal-docs/internal/clients/temporal"
	"temporal-docs/internal/service/schedule"

	"go.uber.org/zap"

	xcloser "github.com/syth0le/gopnik/closer"
)

type App struct {
	Config *configuration.Config
	Logger *zap.Logger
	Closer *xcloser.Closer
}

func New(cfg *configuration.Config, logger *zap.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
		Closer: xcloser.NewCloser(logger, cfg.Application.GracefulShutdownTimeout, cfg.Application.ForceShutdownTimeout, syscall.SIGINT, syscall.SIGTERM),
	}
}

func (a *App) Run() error {
	_, cancelFunction := context.WithCancel(context.Background())
	a.Closer.Add(func() error {
		cancelFunction()
		return nil
	})

	envStruct, err := a.constructEnv()
	if err != nil {
		return fmt.Errorf("construct env: %w", err)
	}

	httpServer := a.newHTTPServer(envStruct)
	a.Closer.Add(httpServer.GracefulStop()...)

	a.Closer.Run(httpServer.Run()...)
	a.Closer.Wait()
	return nil
}

type env struct {
	scheduleService schedule.Service
}

func (a *App) constructEnv() (*env, error) {
	temporalClient, err := a.makeTemporalClient()
	if err != nil {
		return nil, fmt.Errorf("make temporal client: %w", err)
	}

	return &env{
		scheduleService: schedule.NewService(a.Logger, temporalClient),
	}, nil
}

func (a *App) makeTemporalClient() (temporal.Client, error) {
	if !a.Config.Temporal.Enable {
		return temporal.NewClientMock(a.Logger)
	}

	client, err := temporal.NewClient(a.Logger, a.Config.Temporal)
	if err != nil {
		return nil, fmt.Errorf("new temporal client: %w", err)
	}

	a.Closer.Add(client.Close)

	return client, nil
}
