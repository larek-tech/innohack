package view

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/auth/service"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
)

func (v *View) LoginPage(c *fiber.Ctx) error {
	inp := model.LoginReq{}

	return adaptor.HTTPHandler(
		templ.Handler(
			LoginPage(inp),
		),
	)(c)
}

func (v *View) Login(c *fiber.Ctx) error {
	var input model.LoginReq

	if err := c.BodyParser(&input); err != nil {
		v.log.Err(pkg.WrapErr(err)).Msg("parsing login input")
		return adaptor.HTTPHandler(
			templ.Handler(
				LoginForm(input.Email, input.Password, shared.ErrInvalidCredentials),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	token, err := v.service.LoginWithEmail(c.Context(), &service.EmailLoginData{
		Email:    input.Email,
		Password: input.Password,
	}, string(c.Request().Header.UserAgent()))

	if err != nil {
		return adaptor.HTTPHandler(
			templ.Handler(
				LoginForm(input.Email, input.Password, shared.ErrInvalidCredentials),
				templ.WithStatus(fiber.StatusUnauthorized),
			),
		)(c)
	}
	c.Cookie(v.service.CreateAuthCookie(token))
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.SendStatus(fiber.StatusOK)
}
