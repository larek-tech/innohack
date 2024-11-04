package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

func (ctrl *Controller) Login(ctx context.Context, req model.LoginReq) (model.TokenResp, error) {
	user, err := ctrl.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return model.TokenResp{}, err
	}

	token, err := jwt.AuthenticateUser(
		user.ID,
		user.Email,
		user.Password,
		req.Password,
		ctrl.jwtSecret,
	)
	if err != nil {
		return model.TokenResp{}, err
	}
	return token, nil
}
