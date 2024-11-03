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

	session, err := h.service.InsertSession(c.Context(), userID)
	if err != nil {
		return pkg.WrapErr(err)
	}
	return c.Status(fiber.StatusCreated).JSON(session)
}

func (h *Handler) GetSessionContent(c *fiber.Ctx) error {
	sessionID, err := c.ParamsInt("session_id")
	if err != nil {
		return pkg.WrapErr(err)
	}

	content, err := h.service.GetSessionContent(c.Context(), int64(sessionID))
	if err != nil {
		return pkg.WrapErr(err)
	}
	return c.Status(fiber.StatusOK).JSON(content)
}

func (h *Handler) ListSessions(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return pkg.WrapErr(err)
	}

	sessions, err := h.service.ListSessions(c.Context(), userID)
	if err != nil {
		return pkg.WrapErr(err, "list sessions")
	}
	return c.Status(fiber.StatusOK).JSON(sessions)
}

func (h *Handler) UpdateSessionTitle(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return pkg.WrapErr(err)
	}

	sessionID, err := c.ParamsInt("session_id")
	if err != nil {
		return pkg.WrapErr(err)
	}

	title := c.Params("title")
	if err := h.service.UpdateSessionTitle(c.Context(), int64(sessionID), int64(userID), title); err != nil {
		return pkg.WrapErr(err)
	}
	return c.SendStatus(fiber.StatusOK)
}
