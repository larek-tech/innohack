package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (s *Service) InsertQuery(ctx context.Context, query model.QueryDto) (int64, error) {
	queryID, err := s.qr.InsertQuery(ctx, model.Query{
		SessionID: query.SessionID,
		Prompt:    query.Prompt,
	})
	if err != nil {
		return 0, pkg.WrapErr(err, "insert query")
	}
	return queryID, nil
}
