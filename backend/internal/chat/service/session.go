package service

import (
	"context"
	"strconv"

	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

func (s *Service) InsertSession(ctx context.Context, cookie string) (int64, error) {
	token, err := jwt.VerifyAccessToken(cookie, s.jwtSecret)
	if err != nil {
		return 0, pkg.WrapErr(err, "verify token")
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return 0, pkg.WrapErr(err, "get user claims from token")
	}

	userID, err := strconv.ParseFloat(userIDString, 64)
	if err != nil {
		return 0, pkg.WrapErr(err, "parse user id")
	}

	sessionID, err := s.sr.InsertSession(ctx, int64(userID))
	if err != nil {
		return 0, pkg.WrapErr(err, "insert session")
	}

	return sessionID, nil
}
