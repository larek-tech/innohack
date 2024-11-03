package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (h *Handler) InsertSession(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	session, err := h.ctrl.InsertSession(c.Context(), userID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(session)
}
