package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (ctrl *Controller) InsertResponse(ctx context.Context, sessionID int64, resp model.ResponseDto) error {
	if err := ctrl.rr.InsertResponse(ctx, model.Response{
		SessionID:   sessionID,
		QueryID:     resp.QueryID,
		Sources:     resp.Sources,
		Filenames:   resp.Filenames,
		Description: resp.Description,
	}); err != nil {
		return pkg.WrapErr(err, "insert response")
	}
	return nil
}
