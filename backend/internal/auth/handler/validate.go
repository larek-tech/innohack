package handler

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/larek-tech/innohack/backend/internal/auth/view"
	"github.com/larek-tech/innohack/backend/internal/shared"
)

func (h *Handler) ValidateEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if email == "test@test.com" {
		return adaptor.HTTPHandler(templ.Handler(view.EmailField(email, shared.ErrEmailAlreadyTaken)))(c)
	}
	return adaptor.HTTPHandler(templ.Handler(view.EmailField(email, nil)))(c)
}
