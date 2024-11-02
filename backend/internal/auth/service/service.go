package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/auth/repo"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
)

type authRepo interface {
	InsertUser(ctx context.Context, user model.User) (int64, error)
	FindUserByEmail(ctx context.Context, email string) (model.User, error)
}

type Service struct {
	repo      authRepo
	jwtSecret string
	validate  *validator.Validate
}

func New(pg *postgres.Postgres, jwtSecret string) *Service {
	return &Service{
		repo:      repo.New(pg),
		jwtSecret: jwtSecret,
		validate:  validator.New(validator.WithRequiredStructEnabled()),
	}
}
