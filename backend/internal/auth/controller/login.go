package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

func (ctrl *Controller) Login(ctx context.Context, req model.LoginReq) (model.TokenResp, error) {
	user, err := ctrl.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return model.TokenResp{}, pkg.WrapErr(err, "find user by email")
	}

	token, err := jwt.AuthenticateUser(
		user.ID,
		user.Email,
		user.Password,
		req.Password,
		ctrl.jwtSecret,
	)
	if err != nil {
		return model.TokenResp{}, pkg.WrapErr(err, "jwt auth")
	}

	return token, nil
}
