package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

func (ctrl *Controller) Signup(ctx context.Context, req model.SignupReq) (model.TokenResp, error) {
	hashedPass, err := hashPassword(req.Password)
	if err != nil {
		return model.TokenResp{}, err
	}

	userID, err := ctrl.repo.InsertUser(ctx, model.User{
		Email:    req.Email,
		Password: hashedPass,
	})
	if err != nil {
		if pkg.CheckDuplicateKey(err) {
			return model.TokenResp{}, shared.ErrDuplicateKey
		}
		return model.TokenResp{}, shared.ErrStorageInternal
	}

	token, err := jwt.CreateAccessToken(userID, req.Email, ctrl.jwtSecret)
	if err != nil {
		return model.TokenResp{}, err
	}

	return model.TokenResp{
		Token: token,
		Type:  shared.BearerType,
	}, nil
}
