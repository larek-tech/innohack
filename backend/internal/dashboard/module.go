package dashboard

import (
	"github.com/gofiber/fiber/v2"
)

type dashboardHandler interface {
	GetCharts(c *fiber.Ctx) error
}

func InitRoutes(api fiber.Router, h dashboardHandler) {
	dashboard := api.Group("/dashboard")

	dashboard.Post("/", h.GetCharts)
}
