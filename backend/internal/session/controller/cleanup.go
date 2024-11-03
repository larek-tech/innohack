package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (ctrl *Controller) Cleanup(ctx context.Context, sessionID uuid.UUID, userID int64) error {
	session, err := ctrl.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return err
	}

	if session.UserID != userID {
		return shared.ErrNoAccessToSession
	}

	content, err := ctrl.repo.GetSessionContent(ctx, sessionID)
	if err != nil {
		return err
	}

	if len(content) == 0 {
		if err := ctrl.repo.DeleteSession(ctx, sessionID); err != nil {
			return err
		}
	}
	return nil
}
