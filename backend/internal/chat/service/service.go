package service

import (
	"github.com/rs/zerolog"
)

type Service struct {
	log *zerolog.Logger
}

func New(log *zerolog.Logger) *Service {
	return &Service{
		log: log,
	}
}
