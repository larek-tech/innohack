package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

type authService interface {
	Signup(c context.Context, req model.SignupReq) (model.TokenResp, error)
	Login(c context.Context, req model.LoginReq) (model.TokenResp, error)
}

type Handler struct {
	service  authService
	log      *zerolog.Logger
	validate *validator.Validate
	tracer   trace.Tracer
}

func New(tracer trace.Tracer, log *zerolog.Logger, s authService) *Handler {
	return &Handler{
		log:      log,
		service:  s,
		validate: validator.New(validator.WithRequiredStructEnabled()),
		tracer:   tracer,
	}
}
