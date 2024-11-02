package service

import (
	"context"
	"strconv"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

type sessionRepo interface {
	InsertSession(ctx context.Context, userID int64) (int64, error)
	GetSessionContent(ctx context.Context, sessionID int64) ([]model.SessionContent, error)
	ListSessions(ctx context.Context, userID int64) ([]model.Session, error)
	UpdateSessionTitle(ctx context.Context, sessionID int64, title string) error
}

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

func (s *Service) GetSessionContent(ctx context.Context, sessionID int64) ([]*model.SessionContentDto, error) {
	content, err := s.sr.GetSessionContent(ctx, sessionID)
	if err != nil {
		return nil, pkg.WrapErr(err, "get session content")
	}

	res := make([]*model.SessionContentDto, len(content))
	for idx := range len(content) {
		contentDto, err := content[idx].ToDto()
		if err != nil {
			return nil, pkg.WrapErr(err, "dto session content")
		}
		res[idx] = contentDto
	}
	return res, nil
}

func (s *Service) ListSessions(ctx context.Context, userID int64) ([]*model.SessionDto, error) {
	sessions, err := s.sr.ListSessions(ctx, userID)
	if err != nil {
		return nil, pkg.WrapErr(err, "list sessions")
	}

	res := make([]*model.SessionDto, len(sessions))
	for idx := range len(sessions) {
		res[idx] = sessions[idx].ToDto()
	}
	return res, nil
}

func (s *Service) UpdateSessionTitle(ctx context.Context, sessionID int64, title string) error {
	if err := s.sr.UpdateSessionTitle(ctx, sessionID, title); err != nil {
		return pkg.WrapErr(err, "update session title")
	}
	return nil
}
