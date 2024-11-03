package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
)

// Login godoc
//
//	@Summary		Login
//	@Description	Логин пользователя
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.LoginReq	true	"Запрос на логин"
//	@Success		200		{object}	model.TokenResp
//	@Router			/auth/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	ctx, span := h.tracer.Start(c.Context(), "auth.handler.login")
	defer span.End()

	var req model.LoginReq

	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := h.validate.Struct(req); err != nil {
		return err
	}

	token, err := h.ctrl.Login(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(token)
}
