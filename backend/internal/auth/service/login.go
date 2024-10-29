package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/shared"
)

type EmailRegisterData struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
}

// TODO: define service level errors
func (s *Service) RegisterEmail(ctx context.Context, payload *EmailRegisterData, userAgent string) (string, error) {
	// 1. Проверяем что данные валидные
	err := s.validate.Struct(payload)
	if err != nil {
		return "", err
	}
	payload.Password = hashPassword(payload.Password)
	// 2. Создаем запись пользователя
	// 3. Создаем запись о email
	userID, err := s.us.CreateUserWithEmail(ctx, payload)
	if err != nil {
		return "", err
	}
	// 4. Получение токена для авторизации
	sessionToken := createSessionToken(userID, payload.Email)
	// 5*. Отправляем email с кодом подтверждения
	_, err = s.ts.SaveSessionToken(ctx, sessionToken, userAgent, userID)
	if err != nil {
		return "", err
	}
	return sessionToken, nil
}

type EmailLoginData struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// EmailUserData - данные для авторизации через email
type EmailUserData struct {
	UserID   int64
	Email    string
	Password string
}

func (s *Service) LoginWithEmail(ctx context.Context, payload *EmailLoginData, userAgent string) (string, error) {
	err := s.validate.Struct(payload)
	if err != nil {
		return "", shared.ErrInvalidCredentials
	}
	userData, err := s.us.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return "", err
	}
	if err := compareHashAndPassword(userData.Password, payload.Password); err != nil {
		return "", shared.ErrInvalidCredentials
	}

	sessionToken := createSessionToken(userData.UserID, payload.Email)

	_, err = s.ts.SaveSessionToken(ctx, sessionToken, userAgent, userData.UserID)
	if err != nil {
		return "", err
	}
	return sessionToken, nil
}
