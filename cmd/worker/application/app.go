package application

import (
	"context"
	"log"
	"net/http"
	"syscall"

	"temporal-docs/cmd/worker/configuration"
	"temporal-docs/internal/activities"
	"temporal-docs/internal/model"
	"temporal-docs/internal/workflows"

	xcloser "github.com/syth0le/gopnik/closer"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
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

	c, err := client.Dial(client.Options{
		HostPort: a.Config.Temporal.Endpoint,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	a.Closer.AddForce(func() error {
		c.Close()
		return nil
	})

	a.makeTemporalWorker(c)

	a.Closer.Wait()
	return nil
}

func (a *App) makeTemporalWorker(client client.Client) {

	// inject HTTP client into the Activities Struct
	activityManager := &activities.IPActivityManager{
		HTTPClient: http.DefaultClient,
	}

	var w worker.Worker

	// Register Workflow and Activities by worker type
	switch a.Config.Application.Type {
	case model.HotScheduleType:
		// Create the Temporal worker
		w = worker.New(client, workflows.IPAddressAndLocationQueueName, worker.Options{})

		w.RegisterWorkflow(workflows.GetAddressFromIP)
		w.RegisterActivity(activityManager.GetIP)
		w.RegisterActivity(activityManager.GetLocationInfo)
	case model.ColdScheduleType:
		// Create the Temporal worker
		w = worker.New(client, workflows.IPAddressQueueName, worker.Options{})

		w.RegisterWorkflow(workflows.GetOnlyIP)
		w.RegisterActivity(activityManager.GetIP)
	}

	// Start the Worker
	a.Closer.Run(func() error {
		return w.Run(worker.InterruptCh())
	})
	a.Closer.Add(func() error {
		w.Stop()
		return nil
	})
}
