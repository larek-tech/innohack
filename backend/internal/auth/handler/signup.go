package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
)

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
