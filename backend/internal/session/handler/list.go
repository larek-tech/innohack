package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (h *Handler) ListSessions(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	sessions, err := h.ctrl.ListSessions(c.Context(), userID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(sessions)
}
