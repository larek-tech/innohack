package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

// InsertSession godoc
//
//	@Summary		Добавление сессии
//	@Description	Добавление сессии
//	@Tags			session
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	model.Session
//	@Router			/api/session [post]
func (h *Handler) InsertSession(c *fiber.Ctx) error {
	ctx, span := h.tracer.Start(c.Context(), "session.handler.insert_session")
	defer span.End()

	userID, err := strconv.ParseInt(c.Locals(shared.UserIDKey).(string), 10, 64)
	if err != nil {
		return err
	}

	session, err := h.ctrl.InsertSession(ctx, userID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(session)
}
