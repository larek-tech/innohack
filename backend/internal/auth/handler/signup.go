package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/auth/view"
)

func (h *Handler) SignupPage(c *fiber.Ctx) error {
	inp := model.SignupReq{}

	// TODO: return rendered template with sign-up form
	return adaptor.HTTPHandler(
		templ.Handler(
			view.SignupPage(inp),
		),
	)(c)
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	// TODO: receive sing-up information for email auth
	var input model.SignupReq
	err := c.BodyParser(&input)
	if err != nil {
		h.log.Err(err).Msg("unable to parse")
		return adaptor.HTTPHandler(
			templ.Handler(
				view.SignupForm(input.Email, input.Password, input.PasswordConfirm),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}
	// 0. Invalid struct
	err = h.validate.Struct(&input)
	if err != nil {
		h.log.Err(err).Msg("unable to parse")
		return adaptor.HTTPHandler(
			templ.Handler(
				view.SignupForm(input.Email, input.Password, input.PasswordConfirm),
				templ.WithStatus(fiber.StatusAccepted),
			),
		)(c)
	}

	// 1. password mismatch

	if input.Password != input.PasswordConfirm {
		return adaptor.HTTPHandler(
			templ.Handler(
				view.SignupForm(input.Email, input.Password, input.PasswordConfirm),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	// 2. already exists

	if input.Email == "test@test.com" {
		return adaptor.HTTPHandler(
			templ.Handler(
				view.SignupForm(input.Email, input.Password, input.PasswordConfirm),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.JSON(fiber.Map{"hello": "world"})
}
