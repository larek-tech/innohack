package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/handler"
	"github.com/larek-tech/innohack/backend/internal/auth/view"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type authHandler interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	OAuth(c *fiber.Ctx) error
}

type authView interface {
	SignupPage(c *fiber.Ctx) error
	LoginPage(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ValidateEmail(c *fiber.Ctx) error
	OAuth(c *fiber.Ctx) error
}

type AuthModule struct {
	log   *zerolog.Logger
	api   authHandler
	views authView
}

func New() *AuthModule {
	logger := log.With().Str("module", "auth").Logger()
	return &AuthModule{
		log:   &logger,
		api:   handler.New(&logger),
		views: view.New(&logger),
	}
}

func (m *AuthModule) InitRoutes(apiRouter, viewRouter fiber.Router) {
	api := apiRouter.Group("/auth")
	m.initAPI(api)

	views := viewRouter.Group("/auth")
	m.initViews(views)
}

func (m *AuthModule) initAPI(api fiber.Router) {
	api.Post("/signup", m.api.Signup)
	api.Post("/login", m.api.Login)
	api.Post("/oauth", m.api.OAuth)
}

func (m *AuthModule) initViews(views fiber.Router) {
	views.Get("/signup", m.views.SignupPage)
	views.Post("/signup/validate/email", m.views.ValidateEmail)
	views.Get("/login", m.views.LoginPage)
	views.Post("/login", m.views.Login)
	views.Get("/oauth", m.views.OAuth)
}
