package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"go.opentelemetry.io/otel/trace"
)

type chatController interface {
	GetDescription(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64)
	InsertQuery(ctx context.Context, sessionID uuid.UUID, query model.QueryDto) (int64, error)
	InsertResponse(ctx context.Context, sessionID uuid.UUID, resp model.ResponseDto) error
}

type sessionController interface {
	Cleanup(ctx context.Context, sessionID uuid.UUID, userID int64) error
}

type Handler struct {
	cc        chatController
	sc        sessionController
	tracer    trace.Tracer
	jwtSecret string
}

func New(tracer trace.Tracer, jwtSecret string, ctrl chatController) *Handler {
	return &Handler{
		cc:        ctrl,
		tracer:    tracer,
		jwtSecret: jwtSecret,
	}
}
