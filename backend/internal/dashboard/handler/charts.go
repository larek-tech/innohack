package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/dashboard/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (h *Handler) GetCharts(c *fiber.Ctx) error {
	var filter model.Filter

	if err := c.BodyParser(&filter); err != nil {
		return pkg.WrapErr(err)
	}

	report, err := h.ctrl.GetCharts(c.Context(), filter)
	if err != nil {
		return pkg.WrapErr(err)
	}
	return c.Status(fiber.StatusOK).JSON(report)
}
