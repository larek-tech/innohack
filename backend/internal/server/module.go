package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth"
	"github.com/larek-tech/innohack/backend/internal/auth/middleware"
	"github.com/larek-tech/innohack/backend/internal/chat"
)

type module interface {
	InitRoutes(api fiber.Router)
}

func (s *Server) initModules() {
	api := s.app.Group("/api")
	apiWithAuth := api.Use(middleware.Jwt(s.cfg.Server.JwtSecret))

	modules := map[fiber.Router]module{
		api:         auth.New(s.tracer, s.pg, s.cfg.Server.JwtSecret),
		apiWithAuth: chat.New(s.tracer, s.pg, s.cfg.Server.JwtSecret, s.grpcConn.GetConn()),
	}
	for router, module := range modules {
		module.InitRoutes(router)
	}
}
