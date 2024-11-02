package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth"
	"github.com/larek-tech/innohack/backend/internal/chat"
)

type module interface {
	InitRoutes(views fiber.Router)
}

func (s *Server) initModules() {
	views := s.app.Group("/")

	modules := []module{
		auth.New(s.pg, s.cfg.Server.JwtSecret),
		chat.New(),
	}
	for _, m := range modules {
		m.InitRoutes(views)
	}
}
