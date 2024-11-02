package view

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
	"github.com/larek-tech/innohack/backend/internal/chat/service"

	"github.com/rs/zerolog"
)

type chatService interface {
	ProcessMessage(ctx context.Context, req model.Query, out chan<- model.Response, cancel <-chan int64)
}

type View struct {
	service *service.Service
	log     *zerolog.Logger
}

func New(log *zerolog.Logger, s *service.Service) *View {
	return &View{
		service: s,
		log:     log,
	}
}
