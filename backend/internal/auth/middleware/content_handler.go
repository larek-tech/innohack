package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	jsonType = "application/json"
	htmlType = "text/html"
)

func ContentHandler(api, view fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		switch t := http.DetectContentType(c.Request().Header.ContentType()); t {
		case jsonType:
			return api(c)
		case htmlType:
			return view(c)
		default:
			fmt.Println(t)
			return fiber.ErrUnsupportedMediaType
		}
	}
}
