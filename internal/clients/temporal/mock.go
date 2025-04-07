package temporal

import (
	"context"

	"go.uber.org/zap"
)

type ClientMock struct {
	logger *zap.Logger
}

func NewClientMock(logger *zap.Logger) (Client, error) {
	return &ClientMock{logger: logger}, nil
}

func (c *ClientMock) Close() error {
	return nil
}

func (c *ClientMock) CreateSchedule(ctx context.Context, params *CreateScheduleParams) error {
	c.logger.Sugar().Infof("handled create schedule method in temporal client mock: %v", params)

	return nil
}
