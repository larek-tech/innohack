package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
)

func (ctrl *Controller) InsertQuery(ctx context.Context, sessionID uuid.UUID, query model.QueryDto) (int64, error) {
	queryID, err := ctrl.qr.InsertQuery(ctx, model.Query{
		SessionID: sessionID,
		Prompt:    query.Prompt,
	})
	if err != nil {
		return 0, err
	}
	return queryID, nil
}
