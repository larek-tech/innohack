package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

func (s *Service) Login(ctx context.Context, req model.LoginReq) (string, error) {
	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return "", pkg.WrapErr(err, "find user by email")
	}

	if !compareHashAndPassword(user.Password, req.Password) {
		return "", pkg.WrapErr(shared.ErrInvalidCredentials)
	}

	token, err := jwt.CreateAccessToken(user.ID, req.Email, s.jwtSecret)
	if err != nil {
		return "", pkg.WrapErr(err, "create access token")
	}

	return token, nil
}
