package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	expireDelta    = time.Minute * 30
	authCookieName = "name"
)

func (s *Service) CreateAuthCookie(sessionToken string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:    authCookieName,
		Value:   sessionToken,
		Expires: time.Now().Add(expireDelta),
	}
}
