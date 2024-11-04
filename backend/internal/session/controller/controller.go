package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/session/model"
)

type sessionRepo interface {
	InsertSession(ctx context.Context, sessionID uuid.UUID, userID int64) error
	GetSessionByID(ctx context.Context, sessionID uuid.UUID) (model.Session, error)
	GetSessionContent(ctx context.Context, sessionID uuid.UUID) ([]model.SessionContent, error)
	ListSessions(ctx context.Context, userID int64) ([]model.Session, error)
	UpdateSessionTitle(ctx context.Context, sessionID uuid.UUID, userID int64, title string) error
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
}

type Controller struct {
	repo sessionRepo
}

func New(repo sessionRepo) *Controller {
	return &Controller{
		repo: repo,
	}
}
