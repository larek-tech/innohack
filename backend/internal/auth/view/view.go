package view

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/rs/zerolog"
)

type authService interface {
	SignUp(c context.Context, req model.SignUpReq) (string, error)
	Login(c context.Context, req model.LoginReq) (string, error)
}
type View struct {
	// TODO: create interface in auth package
	service  authService
	log      *zerolog.Logger
	validate *validator.Validate
}

func New(log *zerolog.Logger, s authService) *View {
	return &View{
		log:      log,
		service:  s,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *View) authCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:    shared.AuthCookieName,
		Value:   token,
		Expires: time.Now().Add(shared.AuthExp),
	}
}
