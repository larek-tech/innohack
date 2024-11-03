package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

func (s *Service) Signup(ctx context.Context, req model.SignupReq) (model.TokenResp, error) {
	hashedPass, err := hashPassword(req.Password)
	if err != nil {
		return model.TokenResp{}, pkg.WrapErr(err, "generate hash")
	}

	userID, err := s.repo.InsertUser(ctx, model.User{
		Email:    req.Email,
		Password: hashedPass,
	})
	if err != nil {
		if pkg.CheckDuplicateKey(err) {
			return model.TokenResp{}, pkg.WrapErr(shared.ErrDuplicateKey, err.Error())
		}
		return model.TokenResp{}, pkg.WrapErr(shared.ErrStorageInternal, err.Error())
	}

	token, err := jwt.CreateAccessToken(userID, req.Email, s.jwtSecret)
	if err != nil {
		return model.TokenResp{}, pkg.WrapErr(err, "create access token")
	}

	return model.TokenResp{
		Token: token,
		Type:  shared.BearerType,
	}, nil
}
