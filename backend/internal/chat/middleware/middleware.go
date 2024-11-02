package middleware

import (
	"errors"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// WsProtocolUpgrade проверяет, что запрос является вебсокет-апгрейдом.
func WsProtocolUpgrade() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			return ctx.Next()
		}
		return errors.New("popopo")
	}
}
