package service

import (
	"github.com/larek-tech/innohack/backend/internal/chat/repo"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/rs/zerolog"
)

type Service struct {
	jwtSecret string
	sr        sessionRepo
	qr        queryRepo
	rr        responseRepo
	log       *zerolog.Logger
}

func New(log *zerolog.Logger, jwtSecret string, pg *postgres.Postgres) *Service {
	return &Service{
		jwtSecret: jwtSecret,
		log:       log,
		sr:        repo.NewSessionRepo(pg),
		qr:        repo.NewQueryRepo(pg),
		rr:        repo.NewResponseRepo(pg),
	}
}
