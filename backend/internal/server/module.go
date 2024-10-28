package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth"
)

type module interface {
	InitRoutes(api, views fiber.Router)
}

func (s *Server) initModules() {
	api := s.app.Group("/api")
	views := s.app.Group("/")

	modules := []module{
		auth.New(),
	}

	for _, m := range modules {
		m.InitRoutes(api, views)
	}
}
