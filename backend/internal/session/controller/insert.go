package controller

import (
	"context"
	"time"

	"github.com/larek-tech/innohack/backend/internal/session/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (ctrl *Controller) InsertSession(ctx context.Context, userID int64) (model.SessionDto, error) {
	sessionID, err := ctrl.repo.InsertSession(ctx, int64(userID))
	if err != nil {
		return model.SessionDto{}, pkg.WrapErr(err, "insert session")
	}

	return model.SessionDto{
		ID:        sessionID,
		CreatedAt: time.Now(),
	}, nil
}
