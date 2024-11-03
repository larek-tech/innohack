package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (h *Handler) UpdateSessionTitle(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	sessionID, err := c.ParamsInt("session_id")
	if err != nil {
		return err
	}

	title := c.Params("title")
	if err := h.ctrl.UpdateSessionTitle(c.Context(), int64(sessionID), int64(userID), title); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
