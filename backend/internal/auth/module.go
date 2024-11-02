package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/handler"
	"github.com/larek-tech/innohack/backend/internal/auth/service"
	"github.com/larek-tech/innohack/backend/internal/auth/view"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type authHandler interface {
	SignUp(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authView interface {
	SignUpPage(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	LoginPage(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ValidateEmail(c *fiber.Ctx) error
}

type AuthModule struct {
	s     *service.Service
	log   *zerolog.Logger
	api   authHandler
	views authView
}

func New(s *service.Service) *AuthModule {
	logger := log.With().Str("module", "auth").Logger()
	return &AuthModule{
		s:     s,
		log:   &logger,
		api:   handler.New(&logger, s),
		views: view.New(&logger, s),
	}
}

func (m *AuthModule) InitRoutes(apiRouter, viewRouter fiber.Router) {
	api := apiRouter.Group("/auth")
	m.initAPI(api)

	views := viewRouter.Group("/auth")
	m.initViews(views)
}

func (m *AuthModule) initAPI(api fiber.Router) {
	api.Post("/signup", m.api.SignUp)
	api.Post("/login", m.api.Login)
}

func (m *AuthModule) initViews(views fiber.Router) {
	views.Get("/signup", m.views.SignUpPage)
	views.Post("/signup", m.views.SignUp)
	views.Post("/signup/validate/email", m.views.ValidateEmail)
	views.Get("/login", m.views.LoginPage)
	views.Post("/login", m.views.Login)
	// views.Get("/oauth", m.views.OAuth)
}
