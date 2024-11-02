package chat

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/chat/service"
	"github.com/larek-tech/innohack/backend/internal/chat/view"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type chatView interface {
	ChatPage(c *fiber.Ctx) error
	ProcessConn(c *websocket.Conn)
}

type ChatModule struct {
	s     *service.Service
	log   *zerolog.Logger
	views chatView
}

func New(pg *postgres.Postgres, jwtSecret string) *ChatModule {
	logger := log.With().Str("module", "auth").Logger()
	chatService := service.New(&logger, jwtSecret, pg)
	return &ChatModule{
		s:     chatService,
		log:   &logger,
		views: view.New(&logger, chatService),
	}
}

func (m *ChatModule) InitRoutes(viewRouter fiber.Router) {
	views := viewRouter.Group("/chat")
	m.initViews(views)
}

func (m *ChatModule) initViews(views fiber.Router) {
	views.Get("/", m.views.ChatPage)
	views.Get("/ws", websocket.New(m.views.ProcessConn, websocket.Config{
		HandshakeTimeout: 20 * time.Second,
	}), func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
}
