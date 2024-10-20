package server

import "github.com/gofiber/fiber/v2"

type Module interface {
	Name() string
	InitRoutes(*fiber.App) error
}
