package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var req model.LoginReq

	if err := c.BodyParser(&req); err != nil {
		return pkg.WrapErr(err, "unmarshal")
	}

	if err := h.validate.Struct(req); err != nil {
		return pkg.WrapErr(err, "validation")
	}

	token, err := h.service.Login(c.Context(), req)
	if err != nil {
		return pkg.WrapErr(err, "jwt auth")
	}

	return c.Status(fiber.StatusOK).JSON(token)
}
