package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
)

type userStore interface {
	CreateWithEmail(context.Context, *model.EmailRegisterData) (int64, error)
	GetByEmail(context.Context, string) (*model.EmailUserData, error)
}

type tokenStore interface {
	Save(context.Context, string, string, int64) (int64, error)
}

type Service struct {
	users    userStore
	tokens   tokenStore
	validate *validator.Validate
}

func New(us userStore, ts tokenStore) *Service {
	return &Service{
		users:    us,
		tokens:   ts,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
