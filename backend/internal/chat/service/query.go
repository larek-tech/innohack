package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

type queryRepo interface {
	InsertQuery(ctx context.Context, data model.Query) (int64, error)
}

func (s *Service) InsertQuery(ctx context.Context, sessionID int64, query model.QueryDto) (int64, error) {
	queryID, err := s.qr.InsertQuery(ctx, model.Query{
		SessionID: sessionID,
		Prompt:    query.Prompt,
	})
	if err != nil {
		return 0, pkg.WrapErr(err, "insert query")
	}
	return queryID, nil
}
