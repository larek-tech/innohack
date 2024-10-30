package view

import (
	"github.com/go-playground/validator/v10"
	"github.com/larek-tech/innohack/backend/internal/auth/service"
	"github.com/rs/zerolog"
)

type View struct {
	// TODO: create interface in auth package
	service  *service.Service
	log      *zerolog.Logger
	validate *validator.Validate
}

func New(log *zerolog.Logger, s *service.Service) *View {
	return &View{
		log:      log,
		service:  s,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
