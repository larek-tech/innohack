package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/pkg"
)

func (ctrl *Controller) UpdateSessionTitle(ctx context.Context, sessionID, userID int64, title string) error {
	if err := ctrl.repo.UpdateSessionTitle(ctx, sessionID, userID, title); err != nil {
		return pkg.WrapErr(err, "update session title")
	}
	return nil
}
