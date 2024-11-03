package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/session/model"
	"go.opentelemetry.io/otel/trace"
)

type sessionController interface {
	InsertSession(ctx context.Context, userID int64) (model.SessionDto, error)
	GetSessionContent(ctx context.Context, sessionID uuid.UUID, userID int64) ([]*model.SessionContentDto, error)
	ListSessions(ctx context.Context, userID int64) ([]*model.SessionDto, error)
	UpdateSessionTitle(ctx context.Context, sessionID uuid.UUID, userID int64, title string) error
	Cleanup(ctx context.Context, sessionID uuid.UUID, userID int64) error
}

type Handler struct {
	ctrl   sessionController
	tracer trace.Tracer
}

func New(tracer trace.Tracer, ctrl sessionController) *Handler {
	return &Handler{
		ctrl:   ctrl,
		tracer: tracer,
	}
}
