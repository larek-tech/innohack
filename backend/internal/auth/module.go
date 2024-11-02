package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/service"
	"github.com/larek-tech/innohack/backend/internal/auth/view"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type authView interface {
	SignUpPage(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	LoginPage(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ValidateEmail(c *fiber.Ctx) error
}

type AuthModule struct {
	log   *zerolog.Logger
	views authView
}

func New(pg *postgres.Postgres, jwtSecret string) *AuthModule {
	logger := log.With().Str("module", "auth").Logger()
	authService := service.New(pg, jwtSecret)
	return &AuthModule{
		log:   &logger,
		views: view.New(&logger, authService),
	}
}

func (m *AuthModule) InitRoutes(viewRouter fiber.Router) {
	views := viewRouter.Group("/auth")
	m.initViews(views)
}

func (m *AuthModule) initViews(views fiber.Router) {
	views.Get("/signup", m.views.SignUpPage)
	views.Post("/signup", m.views.SignUp)
	views.Post("/signup/validate/email", m.views.ValidateEmail)
	views.Get("/login", m.views.LoginPage)
	views.Post("/login", m.views.Login)
}
