package view

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/internal/chat/service"

	"github.com/rs/zerolog"
)

type chatService interface {
	ProcessMessage(ctx context.Context, req model.QueryDto, out chan<- model.ResponseDto, cancel <-chan int64)
	InsertSession(ctx context.Context, cookie string) (int64, error)
	InsertQuery(ctx context.Context, sessionID int64, query model.QueryDto) (int64, error)
	InsertResponse(ctx context.Context, sessionID int64, resp model.ResponseDto) error
}

type View struct {
	service chatService
	log     *zerolog.Logger
}

func New(log *zerolog.Logger, s *service.Service) *View {
	return &View{
		service: s,
		log:     log,
	}
}
