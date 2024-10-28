package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth"
)

type module interface {
	InitRoutes(app *fiber.App)
}

func (s *Server) initModules() {
	modules := []module{
		auth.New(),
	}

	for _, m := range modules {
		m.InitRoutes(s.app)
	}
}
