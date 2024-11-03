package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/dashboard/model"
)

func (h *Handler) GetCharts(c *fiber.Ctx) error {
	var filter model.Filter

	if err := c.BodyParser(&filter); err != nil {
		return err
	}

	report, err := h.ctrl.GetCharts(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(report)
}
