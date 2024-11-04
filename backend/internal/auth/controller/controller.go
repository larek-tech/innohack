package controller

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
)

type authRepo interface {
	InsertUser(ctx context.Context, user model.User) (int64, error)
	FindUserByEmail(ctx context.Context, email string) (model.User, error)
}

type Controller struct {
	repo      authRepo
	jwtSecret string
	validate  *validator.Validate
}

func New(repo authRepo, jwtSecret string) *Controller {
	return &Controller{
		repo:      repo,
		jwtSecret: jwtSecret,
		validate:  validator.New(validator.WithRequiredStructEnabled()),
	}
}
