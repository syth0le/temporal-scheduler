package schedule

import (
	"context"
	"fmt"

	"temporal-docs/internal/clients/temporal"
	"temporal-docs/internal/model"
	"temporal-docs/internal/workflows"

	xerrors "github.com/syth0le/gopnik/errors"
	"go.uber.org/zap"
)

type Service interface {
	CreateSchedule(ctx context.Context, scheduleType model.ScheduleType) error
}

type ServiceImpl struct {
	logger         *zap.Logger
	temporalClient temporal.Client
}

func NewService(logger *zap.Logger, temporalClient temporal.Client) Service {
	return &ServiceImpl{
		logger:         logger,
		temporalClient: temporalClient,
	}
}

func (s *ServiceImpl) CreateSchedule(ctx context.Context, scheduleType model.ScheduleType) error {
	switch scheduleType {
	case model.HotScheduleType:
		err := s.temporalClient.CreateSchedule(ctx, &temporal.CreateScheduleParams{
			Workflow:  workflows.GetAddressFromIP,
			TaskQueue: workflows.IPAddressAndLocationQueueName,
		})
		if err != nil {
			return fmt.Errorf("create schedule: %w", err)
		}
	case model.ColdScheduleType:
		err := s.temporalClient.CreateSchedule(ctx, &temporal.CreateScheduleParams{
			Workflow:  workflows.GetOnlyIP,
			TaskQueue: workflows.IPAddressQueueName,
		})
		if err != nil {
			return fmt.Errorf("create schedule: %w", err)
		}
	default:
		return xerrors.WrapInternalError(fmt.Errorf("unexpected schedule type: %s", scheduleType))
	}

	return nil
}
