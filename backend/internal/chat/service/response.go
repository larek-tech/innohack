package service

import (
	"context"
	"encoding/json"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

type responseRepo interface {
	InsertResponse(ctx context.Context, data model.Response) error
}

func (s *Service) InsertResponse(ctx context.Context, sessionID int64, resp model.ResponseDto) error {
	charts, err := json.Marshal(resp.Charts)
	if err != nil {
		return pkg.WrapErr(err, "marshal charts")
	}

	multipliers, err := json.Marshal(resp.Multipliers)
	if err != nil {
		return pkg.WrapErr(err, "marshal multipliers")
	}

	if err = s.rr.InsertResponse(ctx, model.Response{
		SessionID:   sessionID,
		QueryID:     resp.QueryID,
		Sources:     resp.Sources,
		Filenames:   resp.Filenames,
		Charts:      charts,
		Description: resp.Description,
		Multipliers: multipliers,
	}); err != nil {
		return pkg.WrapErr(err, "insert response")
	}
	return nil
}
