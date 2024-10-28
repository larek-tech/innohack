package view

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type View struct {
	log      *zerolog.Logger
	validate *validator.Validate
}

func New(log *zerolog.Logger) *View {
	return &View{
		log:      log,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
