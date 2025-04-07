package publicapi

import (
	"fmt"
	"net/http"

	"temporal-docs/internal/model"

	xerrors "github.com/syth0le/gopnik/errors"
)

type createScheduleRequest struct {
	ScheduleType model.ScheduleType `json:"schedule_type"`
}

func (r *createScheduleRequest) validate() error {
	switch r.ScheduleType {
	case model.ColdScheduleType, model.HotScheduleType:
	default:
		return xerrors.WrapValidationError(fmt.Errorf("unexpected schedule type: %s", r.ScheduleType))
	}

	return nil
}

func (h *Handler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	handleRequest := func() error {
		ctx := r.Context()

		request, err := parseJSONRequest[createScheduleRequest](r)
		if err != nil {
			return fmt.Errorf("parse json request: %w", err)
		}

		err = request.validate()
		if err != nil {
			return fmt.Errorf("validate request: %w", err)
		}

		err = h.scheduleService.CreateSchedule(ctx, request.ScheduleType)
		if err != nil {
			return fmt.Errorf("create schedule: %w", err)
		}

		return nil
	}

	err := handleRequest()
	if err != nil {
		h.writeError(r.Context(), w, fmt.Errorf("create schedule: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
