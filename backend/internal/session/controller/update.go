package controller

import (
	"context"

	"github.com/google/uuid"
)

func (ctrl *Controller) UpdateSessionTitle(ctx context.Context, sessionID uuid.UUID, userID int64, title string) error {
	if err := ctrl.repo.UpdateSessionTitle(ctx, sessionID, userID, title); err != nil {
		return err
	}
	return nil
}
