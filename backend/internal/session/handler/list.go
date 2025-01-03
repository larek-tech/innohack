package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

// ListSessions godoc
//
//	@Summary		Получение списка сессий
//	@Description	Получение списка сессий
//	@Tags			session
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Session
//	@Router			/api/session/list [get]
func (h *Handler) ListSessions(c *fiber.Ctx) error {
	ctx, span := h.tracer.Start(c.Context(), "session.handler.list_sessions")
	defer span.End()

	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	sessions, err := h.ctrl.ListSessions(ctx, userID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(sessions)
}
