package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/session/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (ctrl *Controller) GetSessionContent(ctx context.Context, sessionID int64) ([]*model.SessionContentDto, error) {
	content, err := ctrl.repo.GetSessionContent(ctx, sessionID)
	if err != nil {
		return nil, pkg.WrapErr(err, "get session content")
	}

	res := make([]*model.SessionContentDto, len(content))
	for idx := range len(content) {
		contentDto := content[idx].ToDto()
		res[idx] = contentDto
	}
	return res, nil
}
