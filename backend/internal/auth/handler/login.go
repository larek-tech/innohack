package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/auth/view"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (h *Handler) LoginPage(c *fiber.Ctx) error {
	inp := model.LoginReq{}

	return adaptor.HTTPHandler(
		templ.Handler(
			view.LoginPage(inp),
		),
	)(c)
}

func (h *Handler) LoginAPI(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) LoginView(c *fiber.Ctx) error {
	var input model.SignupReq

	if err := c.BodyParser(&input); err != nil {
		h.log.Err(pkg.WrapErr(err)).Msg("parsing login input")
		return adaptor.HTTPHandler(
			templ.Handler(
				view.LoginForm(input.Email, input.Password, shared.ErrInvalidCredentials),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	if input.Email != "test@test.com" || input.Password != "password" {
		return adaptor.HTTPHandler(
			templ.Handler(
				view.LoginForm(input.Email, input.Password, shared.ErrInvalidCredentials),
				templ.WithStatus(fiber.StatusUnauthorized),
			),
		)(c)
	}

	c.Response().Header.Add("Hx-Redirect", "/")
	return c.SendStatus(fiber.StatusOK)
}
