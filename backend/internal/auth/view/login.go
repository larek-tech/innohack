package view

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (v *View) LoginPage(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(
		templ.Handler(
			LoginPage(model.LoginReq{}),
		),
	)(c)
}

func (v *View) Login(c *fiber.Ctx) error {
	var req model.LoginReq

	// invalid data
	if err := c.BodyParser(&req); err != nil {
		v.log.Err(pkg.WrapErr(err)).Msg("parsing login input")
		return adaptor.HTTPHandler(
			templ.Handler(
				LoginForm(req.Email, req.Password, shared.ErrInvalidCredentials),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	token, err := v.service.Login(c.Context(), req)
	if err != nil {
		return adaptor.HTTPHandler(
			templ.Handler(
				LoginForm(req.Email, req.Password, shared.ErrInvalidCredentials),
				templ.WithStatus(fiber.StatusUnauthorized),
			),
		)(c)
	}

	c.Cookie(v.authCookie(token))

	c.Response().Header.Add("Hx-Redirect", "/")
	return c.SendStatus(fiber.StatusOK)
}
