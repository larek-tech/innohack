package controller

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/session/model"
)

func (ctrl *Controller) InsertSession(ctx context.Context, userID int64) (model.SessionDto, error) {
	sessionID := uuid.New()

	err := ctrl.repo.InsertSession(ctx, sessionID, userID)
	if err != nil {
		return model.SessionDto{}, err
	}

	return model.SessionDto{
		ID:        sessionID,
		CreatedAt: time.Now(),
	}, nil
}
