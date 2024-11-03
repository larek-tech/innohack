package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth"
	"github.com/larek-tech/innohack/backend/internal/chat"
)

type module interface {
	InitRoutes(api fiber.Router)
}

func (s *Server) initModules() {
	authRouter := s.app.Group("/auth")
	apiRouter := s.app.Group("/api")

	auth.New(authRouter, s.tracer, s.pg, s.cfg.Server.JwtSecret)
	chat.New(apiRouter, s.tracer, s.pg, s.cfg.Server.JwtSecret, s.grpcConn.GetConn())
}
