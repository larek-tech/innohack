package controller

import (
	"context"
)

func (ctrl *Controller) UpdateSessionTitle(ctx context.Context, sessionID, userID int64, title string) error {
	if err := ctrl.repo.UpdateSessionTitle(ctx, sessionID, userID, title); err != nil {
		return err
	}
	return nil
}
