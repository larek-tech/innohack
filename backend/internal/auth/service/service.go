package service

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type userStore interface {
	CreateUserWithEmail(context.Context, *EmailRegisterData) (int64, error)
	GetUserByEmail(context.Context, string) (*EmailUserData, error)
}

type tokenStore interface {
	SaveSessionToken(context.Context, string, string, int64) (int64, error)
}

type Service struct {
	us       userStore
	ts       tokenStore
	validate *validator.Validate
}

func New(us userStore, ts tokenStore) *Service {
	return &Service{
		us:       us,
		ts:       ts,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
