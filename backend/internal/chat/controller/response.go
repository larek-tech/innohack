package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
)

func (ctrl *Controller) InsertResponse(ctx context.Context, sessionID uuid.UUID, resp model.ResponseDto) error {
	if err := ctrl.rr.InsertResponse(ctx, model.Response{
		SessionID:   sessionID,
		QueryID:     resp.QueryID,
		Sources:     resp.Sources,
		Filenames:   resp.Filenames,
		Description: resp.Description,
	}); err != nil {
		return err
	}
	return nil
}
