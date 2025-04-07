package temporal

import (
	"context"
	"fmt"
	"time"

	"temporal-docs/cmd/starter/configuration"
	"temporal-docs/internal/utils"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.uber.org/zap"
)

type Client interface {
	Close() error
	CreateSchedule(ctx context.Context, params *CreateScheduleParams) error
}

type ClientImpl struct {
	conn   client.Client
	logger *zap.Logger
}

func NewClient(logger *zap.Logger, config configuration.TemporalConfig) (Client, error) {
	logger.Info(config.Endpoint)
	c, err := client.Dial(client.Options{
		HostPort: config.Endpoint,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create client connection: %w", err)
	}

	return &ClientImpl{
		conn:   c,
		logger: logger,
	}, nil
}

func (c *ClientImpl) Close() error {
	c.conn.Close()

	return nil
}

type CreateScheduleParams struct {
	Workflow  any
	TaskQueue string
}

func (c *ClientImpl) CreateSchedule(ctx context.Context, params *CreateScheduleParams) error {
	schedule, err := c.conn.ScheduleClient().Create(
		ctx,
		client.ScheduleOptions{
			ID: utils.GenerateSUID(),
			Action: &client.ScheduleWorkflowAction{
				ID:        utils.GenerateWUID(),
				Workflow:  params.Workflow,
				TaskQueue: params.TaskQueue,
				RetryPolicy: &temporal.RetryPolicy{
					InitialInterval:    time.Second, //amount of time that must elapse before the first retry occurs
					MaximumInterval:    time.Minute, //maximum interval between retries
					BackoffCoefficient: 2,           //how much the retry interval increases
					MaximumAttempts:    2,           // Uncomment this if you want to limit attemptss
				},
			},
			Spec: client.ScheduleSpec{
				Intervals: []client.ScheduleIntervalSpec{
					{
						Every: 30 * time.Second,
					},
				},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("schedule client create: %w", err)
	}

	desc, _ := schedule.Describe(ctx)
	c.logger.Sugar().Infof("created schedule: %v", desc.Info)

	return nil
}
