package view

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
)

func (v *View) SignUpPage(c *fiber.Ctx) error {
	inp := model.SignUpReq{}

	// TODO: return rendered template with sign-up form
	return adaptor.HTTPHandler(
		templ.Handler(
			SignUpPage(inp),
		),
	)(c)
}

func (view *View) SignUp(c *fiber.Ctx) error {
	// TODO: receive sing-up information for email auth
	var input model.SignUpReq
	err := c.BodyParser(&input)
	if err != nil {
		view.log.Err(err).Msg("unable to parse")
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(input),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}
	// 0. Invalid struct
	err = view.validate.Struct(&input)
	if err != nil {
		view.log.Err(err).Msg("unable to parse")
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(input),
				templ.WithStatus(fiber.StatusAccepted),
			),
		)(c)
	}

	// 1. password mismatch

	if input.Password != input.PasswordConfirm {
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(input),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	token, err := view.service.RegisterEmail(c.Context(), &model.EmailRegisterData{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}, string(c.Request().Header.UserAgent()))
	if err != nil {
		return adaptor.HTTPHandler(
			templ.Handler(
				SignUpForm(input),
				templ.WithStatus(fiber.StatusUnprocessableEntity),
			),
		)(c)
	}

	c.Cookie(view.service.CreateAuthCookie(token))
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.JSON(fiber.Map{"hello": "world"})
}
