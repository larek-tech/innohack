package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recovermw "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/internal/auth"
	ac "github.com/larek-tech/innohack/backend/internal/auth/controller"
	ah "github.com/larek-tech/innohack/backend/internal/auth/handler"
	ar "github.com/larek-tech/innohack/backend/internal/auth/repo"
	"github.com/larek-tech/innohack/backend/internal/chat"
	cc "github.com/larek-tech/innohack/backend/internal/chat/controller"
	ch "github.com/larek-tech/innohack/backend/internal/chat/handler"
	cr "github.com/larek-tech/innohack/backend/internal/chat/repo"
	"github.com/larek-tech/innohack/backend/internal/dashboard"
	dc "github.com/larek-tech/innohack/backend/internal/dashboard/controller"
	dh "github.com/larek-tech/innohack/backend/internal/dashboard/handler"
	"github.com/larek-tech/innohack/backend/internal/session"
	sc "github.com/larek-tech/innohack/backend/internal/session/controller"
	sh "github.com/larek-tech/innohack/backend/internal/session/handler"
	sr "github.com/larek-tech/innohack/backend/internal/session/repo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/larek-tech/innohack/backend/config"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/grpc_client"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/larek-tech/innohack/backend/pkg/tracing"
	"github.com/rs/zerolog/log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	srvHeader   = "larek.tech"
	bodyLimitMb = 5
)

type Server struct {
	app      *fiber.App
	cfg      config.Config
	pg       *postgres.Postgres
	tracer   trace.Tracer
	exporter sdktrace.SpanExporter
	grpcConn grpc_client.GrpcClient
}

func New(cfg config.Config) Server {
	if err := cfg.Server.Validate(); err != nil {
		panic(pkg.WrapErr(err, "config validation"))
	}

	app := fiber.New(fiber.Config{
		ServerHeader: srvHeader,
		BodyLimit:    bodyLimitMb * 1024 * 1024,
		ErrorHandler: NewErrorHandler(errMap).Handler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Server.GetOrigins(),
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Use(recovermw.New())

	app.Get("/health", healtCheckHandler(uuid.NewString()))

	exporter := tracing.MustNewExporter(context.Background(), cfg.Jaeger.URL())
	provider := tracing.MustNewTraceProvider(exporter, "chat")
	otel.SetTracerProvider(provider)

	tracer := otel.Tracer("chat")

	s := Server{
		app:      app,
		cfg:      cfg,
		pg:       postgres.MustNew(cfg.Postgres, tracer),
		tracer:   tracer,
		exporter: exporter,
		grpcConn: grpc_client.MustNewGrpcClientWithInsecure(cfg.Analytics),
	}

	return s
}

func (s *Server) Serve() {
	defer func() {
		if err := s.exporter.Shutdown(context.Background()); err != nil {
			log.Err(err).Msg("shutodwn exporter")
		}
	}()
	defer s.pg.Close()

	api := s.app.Group("/api")

	authRepo := ar.New(s.pg)
	authCtrl := ac.New(authRepo, s.cfg.Server.JwtSecret)
	authHandler := ah.New(s.tracer, authCtrl)
	auth.InitRoutes(s.app, authHandler)

	queryRepo := cr.NewQueryRepo(s.pg)
	respRepo := cr.NewResponseRepo(s.pg)
	chatCtrl := cc.New(s.cfg.Server.JwtSecret, s.grpcConn.GetConn(), queryRepo, respRepo)
	chatHandler := ch.New(s.tracer, s.cfg.Server.JwtSecret, chatCtrl)
	chat.InitRoutes(api, chatHandler, s.cfg.Server.JwtSecret)

	sessionRepo := sr.New(s.pg)
	sessionCtrl := sc.New(sessionRepo)
	sessionHandler := sh.New(s.tracer, sessionCtrl)
	session.InitRoutes(api, sessionHandler, s.cfg.Server.JwtSecret)

	dashboardCtrl := dc.New(s.grpcConn.GetConn())
	dashboardHandler := dh.New(dashboardCtrl)
	dashboard.InitRoutes(api, dashboardHandler)

	go s.listenHTTP(strconv.Itoa(s.cfg.Server.Port))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	if err := s.app.Shutdown(); err != nil {
		log.Err(pkg.WrapErr(err)).Msg("graceful shutdown")
	}
}

func (s *Server) listenHTTP(port string) {
	addr := net.JoinHostPort("0.0.0.0", port)
	if err := s.app.Listen(addr); err != nil {
		log.Err(pkg.WrapErr(err)).Msg("application interrupted")
	}
}
