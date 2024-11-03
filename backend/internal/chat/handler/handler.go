package handler

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"go.opentelemetry.io/otel/trace"
)

type chatController interface {
	GetDescription(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64)
	InsertQuery(ctx context.Context, sessionID int64, query model.QueryDto) (int64, error)
	InsertResponse(ctx context.Context, sessionID int64, resp model.ResponseDto) error
}

type Handler struct {
	ctrl      chatController
	tracer    trace.Tracer
	jwtSecret string
}

func New(tracer trace.Tracer, jwtSecret string, ctrl chatController) *Handler {
	return &Handler{
		ctrl:      ctrl,
		tracer:    tracer,
		jwtSecret: jwtSecret,
	}
}
