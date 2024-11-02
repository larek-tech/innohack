package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

func (s *Service) SignUp(ctx context.Context, req model.SignUpReq) (string, error) {
	hashedPass, err := hashPassword(req.Password)
	if err != nil {
		return "", pkg.WrapErr(err)
	}

	userID, err := s.repo.InsertUser(ctx, model.User{
		Email:    req.Email,
		Password: hashedPass,
	})
	if err != nil {
		if pkg.CheckDuplicateKey(err) {
			return "", pkg.WrapErr(shared.ErrDuplicateKey, err.Error())
		}
		return "", pkg.WrapErr(shared.ErrStorageInternal, err.Error())
	}

	token, err := jwt.CreateAccessToken(userID, req.Email, s.jwtSecret)
	if err != nil {
		return "", pkg.WrapErr(err, "create access token")
	}

	return token, nil
}
