package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/middleware"
	"go.opentelemetry.io/otel/trace"
)

type sessionHandler interface {
	InsertSession(c *fiber.Ctx) error
	GetSessionContent(c *fiber.Ctx) error
	ListSessions(c *fiber.Ctx) error
	UpdateSessionTitle(c *fiber.Ctx) error
}

func InitRoutes(api fiber.Router, h sessionHandler, secret string, tracer trace.Tracer) {
	session := api.Group("/session")
	session.Use(middleware.Jwt(secret, tracer))

	session.Post("/", h.InsertSession)
	session.Get("/list", h.ListSessions)
	session.Get("/:session_id", h.GetSessionContent)
	session.Put("/:session_id/:title", h.UpdateSessionTitle)
}
