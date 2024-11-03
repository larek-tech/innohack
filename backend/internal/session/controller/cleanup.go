package controller

import (
	"context"
)

func (ctrl *Controller) CleanupSession(ctx context.Context, sessionID int64, token string) error {
	session, err :=

		ctrl.repo.DeleteSession()
}
