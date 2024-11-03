package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/session/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (ctrl *Controller) GetSessionContent(ctx context.Context, sessionID uuid.UUID, userID int64) ([]*model.SessionContentDto, error) {
	content, err := ctrl.repo.GetSessionContent(ctx, sessionID, userID)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, shared.ErrNoAccessToSession
	}

	res := make([]*model.SessionContentDto, len(content))
	for idx := range len(content) {
		contentDto := content[idx].ToDto()
		res[idx] = contentDto
	}
	return res, nil
}
