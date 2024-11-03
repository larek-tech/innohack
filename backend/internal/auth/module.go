package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/handler"
	"github.com/larek-tech/innohack/backend/internal/auth/service"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type authHandler interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type AuthModule struct {
	log     *zerolog.Logger
	handler authHandler
}

func New(tracer trace.Tracer, pg *postgres.Postgres, jwtSecret string) *AuthModule {
	logger := log.With().Str("module", "auth").Logger()
	authService := service.New(pg, jwtSecret)
	return &AuthModule{
		log:     &logger,
		handler: handler.New(tracer, &logger, authService),
	}
}

func (m *AuthModule) InitRoutes(api fiber.Router) {
	auth := api.Group("/")
	auth.Post("/signup", m.handler.Signup)
	auth.Post("/login", m.handler.Login)
}
