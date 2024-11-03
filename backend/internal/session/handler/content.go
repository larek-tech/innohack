package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (h *Handler) GetSessionContent(c *fiber.Ctx) error {
	sessionID, err := c.ParamsInt("session_id")
	if err != nil {
		return pkg.WrapErr(err)
	}

	content, err := h.ctrl.GetSessionContent(c.Context(), int64(sessionID))
	if err != nil {
		return pkg.WrapErr(err)
	}
	return c.Status(fiber.StatusOK).JSON(content)
}
