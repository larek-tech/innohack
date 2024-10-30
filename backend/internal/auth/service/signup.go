package service

import (
	"context"

	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (s *Service) RegisterEmail(ctx context.Context, payload *model.EmailRegisterData, userAgent string) (string, error) {
	// 1. Проверяем что данные валидные
	err := s.validate.Struct(payload)
	if err != nil {
		return "", pkg.WrapErr(err)
	}
	payload.Password = hashPassword(payload.Password)
	// 2. Создаем запись пользователя
	// 3. Создаем запись о email
	userID, err := s.users.CreateWithEmail(ctx, payload)
	if err != nil {
		return "", shared.ErrStorageInternal
	}
	// 4. Получение токена для авторизации
	sessionToken := createSessionToken(userID, payload.Email)
	// 5*. Отправляем email с кодом подтверждения
	_, err = s.tokens.Save(ctx, sessionToken, userAgent, userID)
	if err != nil {
		return "", shared.ErrStorageInternal
	}
	return sessionToken, nil
}
