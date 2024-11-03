package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/session/model"
)

func (ctrl *Controller) ListSessions(ctx context.Context, userID int64) ([]*model.SessionDto, error) {
	sessions, err := ctrl.repo.ListSessions(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]*model.SessionDto, len(sessions))
	for idx := range len(sessions) {
		res[idx] = sessions[idx].ToDto()
	}
	return res, nil
}
