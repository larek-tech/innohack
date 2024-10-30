package server

import (
	"errors"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/server/view"
	"github.com/larek-tech/innohack/backend/templ/pages"
)

var (
	errorHandler = func(c *fiber.Ctx, e error) error {
		var fiberErr *fiber.Error

		// 404 not found
		if errors.As(e, &fiberErr) && fiberErr.Code == fiber.StatusNotFound {
			err := adaptor.HTTPHandler(
				templ.Handler(
					pages.NotFound(c.Path()),
					templ.WithStatus(fiber.StatusNotFound),
				),
			)(c)
			return err
		}

		return e
	}

	indexHanlder = func(c *fiber.Ctx) error {
		index := view.Index()
		return adaptor.HTTPHandler(templ.Handler(index))(c)
	}

	healtCheckHandler = func(nodeID string) fiber.Handler {
		return func(c *fiber.Ctx) error {
			return c.JSON(healthCheckInfo{nodeID})
		}
	}
)

type healthCheckInfo struct {
	NodeID string `json:"node_id"`
}
