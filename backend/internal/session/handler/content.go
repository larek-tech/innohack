package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (h *Handler) GetSessionContent(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	sessionID, err := uuid.Parse(c.Params("session_id"))
	if err != nil {
		return err
	}

	content, err := h.ctrl.GetSessionContent(c.Context(), sessionID, userID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(content)
}
