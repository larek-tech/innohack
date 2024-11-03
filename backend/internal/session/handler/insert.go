package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (h *Handler) InsertSession(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return pkg.WrapErr(err)
	}

	session, err := h.ctrl.InsertSession(c.Context(), userID)
	if err != nil {
		return pkg.WrapErr(err)
	}
	return c.Status(fiber.StatusCreated).JSON(session)
}
