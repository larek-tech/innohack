package server

import (
	"github.com/gofiber/fiber/v2"
)

type module interface {
	InitRoutes(api, views fiber.Router)
}

func (s *Server) initModules(modules ...module) {
	api := s.app.Group("/api")
	views := s.app.Group("/")

	for _, m := range modules {
		m.InitRoutes(api, views)
	}
}
