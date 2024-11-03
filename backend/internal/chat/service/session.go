package service

import (
	"context"
	"time"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

type sessionRepo interface {
	InsertSession(ctx context.Context, userID int64) (int64, error)
	GetSessionContent(ctx context.Context, sessionID int64) ([]model.SessionContent, error)
	ListSessions(ctx context.Context, userID int64) ([]model.Session, error)
	UpdateSessionTitle(ctx context.Context, sessionID, userID int64, title string) error
}

func (s *Service) InsertSession(ctx context.Context, userID int64) (model.SessionDto, error) {
	sessionID, err := s.sr.InsertSession(ctx, int64(userID))
	if err != nil {
		return model.SessionDto{}, pkg.WrapErr(err, "insert session")
	}

	return model.SessionDto{
		ID:        sessionID,
		CreatedAt: time.Now(),
	}, nil
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

func (s *Service) UpdateSessionTitle(ctx context.Context, sessionID, userID int64, title string) error {
	if err := s.sr.UpdateSessionTitle(ctx, sessionID, userID, title); err != nil {
		return pkg.WrapErr(err, "update session title")
	}
	return nil
}
