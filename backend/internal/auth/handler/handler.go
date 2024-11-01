package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/larek-tech/innohack/backend/internal/auth/service"
	"github.com/rs/zerolog"
)

type MailService interface {
	CheckEmail(ctx context.Context, email string) (bool, error)
}

type Handler struct {
	log      *zerolog.Logger
	validate *validator.Validate
	service  *service.Service
}

func New(log *zerolog.Logger, service *service.Service) *Handler {
	return &Handler{
		log:      log,
		validate: validator.New(validator.WithRequiredStructEnabled()),
		service:  service,
	}
}
