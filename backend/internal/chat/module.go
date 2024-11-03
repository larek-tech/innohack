package chat

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/chat/middleware"
)

type chatHandler interface {
	ProcessConn(c *websocket.Conn)
}

func InitRoutes(api fiber.Router, h chatHandler, secret string) {
	chat := api.Group("/chat")

	ws := chat.Group("/ws")
	ws.Use(middleware.WsProtocolUpgrade())
	ws.Get("/:session_id", websocket.New(
		h.ProcessConn,
		websocket.Config{HandshakeTimeout: 20 * time.Second},
	))
}
