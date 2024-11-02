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

func New(s *service.Service) *ChatModule {
	logger := log.With().Str("module", "auth").Logger()
	return &ChatModule{
		s:   s,
		log: &logger,
		// api:   handler.New(&logger, s),
		views: view.New(&logger, s),
	}
}

func (m *ChatModule) InitRoutes(apiRouter, viewRouter fiber.Router) {
	api := apiRouter.Group("/chat")
	m.initAPI(api)

	views := viewRouter.Group("/chat")
	m.initViews(views)
}

func (m *ChatModule) initAPI(api fiber.Router) {
	// api.Post("/signup", m.api.SignUp)
	// api.Post("/login", m.api.Login)
	// api.Get("/oauth", m.api.OAuth)
}

func (m *ChatModule) initViews(views fiber.Router) {
	// views.Use(middleware.WsProtocolUpgrade())
	views.Get("/", m.views.ChatPage)
	views.Get("/ws", websocket.New(m.views.Message, websocket.Config{
		HandshakeTimeout: 20 * time.Second,
	}))
}
