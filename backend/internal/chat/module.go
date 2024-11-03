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

	chat.Use(middleware.WsProtocolUpgrade())
	chat.Get("/ws", websocket.New(
		m.handler.ProcessConn,
		websocket.Config{HandshakeTimeout: 20 * time.Second},
	))
}
