package auth

import (
	"context"

	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type MailService interface {
	CheckEmail(context.Context, string) (bool, error)
	CreateNew(context.Context, ...string)
}

type AuthModule struct {
	mail     MailService
	log      zerolog.Logger
	validate *validator.Validate
}

func NewAuthModule() *AuthModule {
	return &AuthModule{
		log:      log.With().Str("module", "auth").Logger(),
		validate: validator.New(),
	}
}

func (am *AuthModule) Name() string {
	return "auth"
}

func (am *AuthModule) InitRoutes(app *fiber.App) error {
	auth := app.Group("/auth")

	auth.Get("/sing-up", func(c *fiber.Ctx) error {
		inp := RegistrationInput{}

		// TODO: return rendered template with sign-up form
		return adaptor.HTTPHandler(
			templ.Handler(
				RegisterPage(inp),
			),
		)(c)
	})

	auth.Post("/sing-up", func(c *fiber.Ctx) error {
		// TODO: receive sing-up information for email auth
		var input RegistrationInput
		err := c.BodyParser(&input)
		if err != nil {
			am.log.Err(err).Msg("unable to parse")
			return adaptor.HTTPHandler(
				templ.Handler(
					RegisterForm(input.Email, input.Password, input.PasswordConfirm),
					templ.WithStatus(fiber.StatusUnprocessableEntity),
				),
			)(c)
		}
		// 0. Invalid struct
		err = am.validate.Struct(&input)
		if err != nil {
			am.log.Err(err).Msg("unable to parse")
			return adaptor.HTTPHandler(
				templ.Handler(
					RegisterForm(input.Email, input.Password, input.PasswordConfirm),
					templ.WithStatus(fiber.StatusAccepted),
				),
			)(c)
		}

		// 1. password mismatch

		if input.Password != input.PasswordConfirm {
			return adaptor.HTTPHandler(
				templ.Handler(
					RegisterForm(input.Email, input.Password, input.PasswordConfirm),
					templ.WithStatus(fiber.StatusUnprocessableEntity),
				),
			)(c)
		}

		// 2. already exists

		if input.Email == "test@test.com" {
			return adaptor.HTTPHandler(
				templ.Handler(
					RegisterForm(input.Email, input.Password, input.PasswordConfirm),
					templ.WithStatus(fiber.StatusUnprocessableEntity),
				),
			)(c)
		}
		c.Response().Header.Add("Hx-Redirect", "/")
		return c.JSON(fiber.Map{"hello": "world"})
	})

	auth.Post("/sign-up/validate/email", func(c *fiber.Ctx) error {

		email := c.FormValue("email")
		if email == "test@test.com" {
			return adaptor.HTTPHandler(templ.Handler(EmailField(email, ErrEmailAlreadyTaken)))(c)
		}
		return adaptor.HTTPHandler(templ.Handler(EmailField(email, nil)))(c)
	})

	auth.Get("/oauth", func(c *fiber.Ctx) error {
		// TODO: connect multiple oauth via github / yandex id
		return nil
	})

	auth.Get("/login", func(c *fiber.Ctx) error {
		inp := LoginInput{}
		return adaptor.HTTPHandler(
			templ.Handler(
				LoginPage(inp),
			),
		)(c)

	})

	auth.Post("/login", func(c *fiber.Ctx) error {
		var input RegistrationInput
		err := c.BodyParser(&input)
		if err != nil {
			am.log.Err(err).Msg("unable to parse")
			return adaptor.HTTPHandler(
				templ.Handler(
					LoginForm(input.Email, input.Password, ErrInvalidCredentials),
					templ.WithStatus(fiber.StatusUnprocessableEntity),
				),
			)(c)
		}

		if input.Email != "test@test.com" && input.Password != "password" {
			return adaptor.HTTPHandler(
				templ.Handler(
					LoginForm(input.Email, input.Password, ErrInvalidCredentials),
					templ.WithStatus(fiber.StatusUnauthorized),
				),
			)(c)
		}

		c.Response().Header.Add("Hx-Redirect", "/")
		return c.JSON(fiber.Map{"hello": "world"})
	})

	return nil
}
