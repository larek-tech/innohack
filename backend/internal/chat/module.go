package chat

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/larek-tech/innohack/backend/internal/analytics/pb"
	"github.com/larek-tech/innohack/backend/internal/chat/handler"
	"github.com/larek-tech/innohack/backend/internal/chat/middleware"
	"github.com/larek-tech/innohack/backend/internal/chat/service"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type chatHandler interface {
	ProcessConn(c *websocket.Conn)
	InsertSession(c *fiber.Ctx) error
	GetSessionContent(c *fiber.Ctx) error
	ListSessions(c *fiber.Ctx) error
	UpdateSessionTitle(c *fiber.Ctx) error
}

type ChatModule struct {
	s       *service.Service
	log     *zerolog.Logger
	handler chatHandler
}

func New(tracer trace.Tracer, pg *postgres.Postgres, jwtSecret string, grpcConn *grpc.ClientConn) *ChatModule {
	logger := log.With().Str("module", "auth").Logger()

	analytics := pb.NewAnalyticsClient(grpcConn)
	chatService := service.New(&logger, jwtSecret, pg, analytics)

	return &ChatModule{
		s:       chatService,
		log:     &logger,
		handler: handler.New(tracer, &logger, chatService),
	}
}

func (m *ChatModule) InitRoutes(api fiber.Router) {
	chat := api.Group("/chat")

	session := chat.Group("/session")
	session.Post("/", m.handler.InsertSession)
	session.Get("/:session_id", m.handler.GetSessionContent)
	session.Get("/list", m.handler.ListSessions)
	session.Put("/:title", m.handler.UpdateSessionTitle)

	ws := chat.Group("/ws")
	ws.Use(middleware.WsProtocolUpgrade())
	ws.Get("/ws/:session_id", websocket.New(
		m.handler.ProcessConn,
		websocket.Config{HandshakeTimeout: 20 * time.Second},
	))
}
