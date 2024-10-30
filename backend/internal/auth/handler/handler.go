package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type MailService interface {
	CheckEmail(ctx context.Context, email string) (bool, error)
}

type Handler struct {
	log      *zerolog.Logger
	validate *validator.Validate
}

func New(log *zerolog.Logger) *Handler {
	return &Handler{
		log:      log,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
