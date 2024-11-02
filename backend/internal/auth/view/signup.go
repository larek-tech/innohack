package view

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
)

func (v *View) SignUpPage(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(
		templ.Handler(
			SignUpPage(model.SignUpReq{}),
		),
	)(c)
}

func (v *View) SignUp(c *fiber.Ctx) error {
	var req model.SignUpReq

	// check for ivalid input
	if err := c.BodyParser(&req); err != nil {
		v.log.Err(err).Msg("unable to parse")
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(req),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	// validate input
	if err := v.validate.Struct(&req); err != nil {
		v.log.Err(err).Msg("unable to parse")
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(req),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	// password mismatch
	if req.Password != req.PasswordConfirm {
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(req),
				templ.WithStatus(fiber.StatusBadRequest),
			),
		)(c)
	}

	// can't auth via jwt
	token, err := v.service.SignUp(c.Context(), req)
	if err != nil {
		v.log.Err(err).Msg("jwt token")
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(req),
				templ.WithStatus(fiber.StatusInternalServerError),
			),
		)(c)
	}

	c.Cookie(v.authCookie(token))

	c.Response().Header.Add("Hx-Redirect", "/")
	return c.SendStatus(fiber.StatusOK)
}
