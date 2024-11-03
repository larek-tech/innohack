package auth

import (
	"github.com/gofiber/fiber/v2"
)

type authHandler interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

func InitRoutes(app fiber.Router, h authHandler) {
	auth := app.Group("/auth")
	auth.Post("/signup", h.Signup)
	auth.Post("/login", h.Login)
}
