package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

// GetSessionContent godoc
//
//	@Summary		Получение контента сессии
//	@Description	Получение контента сессии
//	@Tags			session
//	@Accept			json
//	@Produce		json
//	@Param			session_id	path		string	true	"ID сессии в формате UUID"
//	@Success		200			{object}	[]model.SessionContentDto
//	@Router			/api/session/{session_id} [get]
func (h *Handler) GetSessionContent(c *fiber.Ctx) error {
	ctx, span := h.tracer.Start(c.Context(), "session.handler.get_session_content")
	defer span.End()

	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	sessionID, err := uuid.Parse(c.Params("session_id"))
	if err != nil {
		return err
	}

	content, err := h.ctrl.GetSessionContent(ctx, sessionID, userID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(content)
}
