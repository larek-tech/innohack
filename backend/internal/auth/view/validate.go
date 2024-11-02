package view

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (v *View) ValidateEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if email == "test@test.com" {
		return adaptor.HTTPHandler(templ.Handler(EmailField(email, shared.ErrEmailAlreadyTaken)))(c)
	}
	return adaptor.HTTPHandler(templ.Handler(EmailField(email, nil)))(c)
}
