package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/chat/model"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) ProcessMessage(ctx context.Context, msg string) (model.Response, error) {

	return model.Response{}, nil
}
