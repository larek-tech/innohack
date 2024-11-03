package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (h *Handler) UpdateSessionTitle(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	sessionID, err := uuid.Parse(c.Params("session_id"))
	if err != nil {
		return err
	}

	title := c.Params("title")
	if err := h.ctrl.UpdateSessionTitle(c.Context(), sessionID, int64(userID), title); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
