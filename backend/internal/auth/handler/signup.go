package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
)

// Signup godoc
//
// @Summary		Signup
// @Description	Регистрация пользователя
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			body	body		model.SignupReq	true	"Запрос на регистрацию"
// @Success		201		{object}	model.TokenResp
// @Router			/auth/signup [post]
func (h *Handler) Signup(c *fiber.Ctx) error {
	var req model.SignupReq

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(&req); err != nil {
		return err
	}

	token, err := h.ctrl.Signup(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(token)
}
