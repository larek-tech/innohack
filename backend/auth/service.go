package auth

import "context"

type UserStore interface {
}

type TokenStore interface {
}

type Service struct {
	userStore  UserStore
	tokenStore TokenStore
}

func NewService(userStore UserStore, tokenStore TokenStore) *Service {
	return &Service{
		userStore:  userStore,
		tokenStore: tokenStore,
	}
}

func (s *Service) Register(ctx context.Context, payload RegistrationInput) error {
	return nil
}
