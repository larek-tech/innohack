package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/handler"
	"github.com/larek-tech/innohack/backend/internal/auth/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type authHandler interface {
	Signup(c *fiber.Ctx) error
	SignupPage(c *fiber.Ctx) error
	ValidateEmail(c *fiber.Ctx) error
	OAuth(c *fiber.Ctx) error
	LoginAPI(c *fiber.Ctx) error
	LoginView(c *fiber.Ctx) error
	LoginPage(c *fiber.Ctx) error
}

type AuthModule struct {
	log     *zerolog.Logger
	handler authHandler
}

func New() *AuthModule {
	logger := log.With().Str("module", "auth").Logger()
	return &AuthModule{
		log:     &logger,
		handler: handler.New(&logger),
	}
}

func (m *AuthModule) InitRoutes(app *fiber.App) {
	g := app.Group("/auth")

	g.Get("/signup", m.handler.SignupPage)
	g.Get("/login", m.handler.LoginPage)
	g.Get("/oauth", m.handler.OAuth)
	g.Post("/signup", m.handler.Signup)
	g.Post("/signup/validate/email", m.handler.ValidateEmail)
	g.Post("/login", middleware.ContentHandler(m.handler.LoginAPI, m.handler.LoginView))
}
