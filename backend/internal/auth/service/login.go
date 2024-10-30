package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (s *Service) LoginWithEmail(ctx context.Context, payload *model.EmailLoginData, userAgent string) (string, error) {
	err := s.validate.Struct(payload)
	if err != nil {
		return "", shared.ErrInvalidCredentials
	}
	userData, err := s.users.GetByEmail(ctx, payload.Email)
	if err != nil {
		return "", shared.ErrStorageInternal
	}
	if err = compareHashAndPassword(userData.Password, payload.Password); err != nil {
		return "", shared.ErrInvalidCredentials
	}

	sessionToken := createSessionToken(userData.UserID, payload.Email)

	_, err = s.tokens.Save(ctx, sessionToken, userAgent, userData.UserID)
	if err != nil {
		return "", shared.ErrStorageInternal
	}
	return sessionToken, nil
}
