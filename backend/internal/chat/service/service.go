package service

import (
	"context"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) ProcessMessage(ctx context.Context, msg string) (string, error) {

	return msg + " " + msg, nil
}
