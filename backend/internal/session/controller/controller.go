package controller

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/session/model"
)

type sessionRepo interface {
	InsertSession(ctx context.Context, userID int64) (int64, error)
	GetSessionContent(ctx context.Context, sessionID, userID int64) ([]model.SessionContent, error)
	ListSessions(ctx context.Context, userID int64) ([]model.Session, error)
	UpdateSessionTitle(ctx context.Context, sessionID, userID int64, title string) error
	DeleteSession(ctx context.Context, sessionID, userID int64) error
}

type Controller struct {
	repo sessionRepo
}

func New(repo sessionRepo) *Controller {
	return &Controller{
		repo: repo,
	}
}
