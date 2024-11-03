package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/middleware"
)

type sessionHandler interface {
	InsertSession(c *fiber.Ctx) error
	GetSessionContent(c *fiber.Ctx) error
	ListSessions(c *fiber.Ctx) error
	UpdateSessionTitle(c *fiber.Ctx) error
}

func InitRoutes(api fiber.Router, h sessionHandler, secret string) {
	session := api.Group("/session")
	session.Use(middleware.Jwt(secret))
	session.Post("/", h.InsertSession)
	session.Get("/:session_id", h.GetSessionContent)
	session.Get("/list", h.ListSessions)
	session.Put("/:title", h.UpdateSessionTitle)
}
