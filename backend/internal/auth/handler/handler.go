package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"go.opentelemetry.io/otel/trace"
)

type authController interface {
	Signup(c context.Context, req model.SignupReq) (model.TokenResp, error)
	Login(c context.Context, req model.LoginReq) (model.TokenResp, error)
}

type Handler struct {
	ctrl     authController
	validate *validator.Validate
	tracer   trace.Tracer
}

func New(tracer trace.Tracer, ctrl authController) *Handler {
	return &Handler{
		ctrl:     ctrl,
		validate: validator.New(validator.WithRequiredStructEnabled()),
		tracer:   tracer,
	}
}
