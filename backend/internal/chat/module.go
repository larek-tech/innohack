package chat

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/chat/service"
	"github.com/larek-tech/innohack/backend/internal/chat/view"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type chatView interface {
	ChatPage(c *fiber.Ctx) error
	Message(c *websocket.Conn)
}

type ChatModule struct {
	s     *service.Service
	log   *zerolog.Logger
	views chatView
}

func New() *ChatModule {
	logger := log.With().Str("module", "auth").Logger()
	chatService := service.New()
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
	// views.Use(middleware.WsProtocolUpgrade())
	views.Get("/", m.views.ChatPage)
	views.Get("/ws", websocket.New(m.views.Message, websocket.Config{
		HandshakeTimeout: 20 * time.Second,
	}))
}
