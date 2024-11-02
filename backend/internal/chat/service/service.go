package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/internal/chat/repo"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/rs/zerolog"
)

type sessionRepo interface {
	InsertSession(ctx context.Context, userID int64) (int64, error)
}

type queryRepo interface {
	InsertQuery(ctx context.Context, data model.Query) (int64, error)
}

type Service struct {
	jwtSecret string
	sr        sessionRepo
	qr        queryRepo
	log       *zerolog.Logger
}

func New(log *zerolog.Logger, jwtSecret string, pg *postgres.Postgres) *Service {
	return &Service{
		jwtSecret: jwtSecret,
		log:       log,
		sr:        repo.NewSessionRepo(pg),
		qr:        repo.NewQueryRepo(pg),
	}
}
