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

	"github.com/larek-tech/innohack/backend/config"
	"github.com/larek-tech/innohack/backend/internal/auth"
	"github.com/larek-tech/innohack/backend/internal/chat"
	authService "github.com/larek-tech/innohack/backend/internal/auth/service"
	chatService "github.com/larek-tech/innohack/backend/internal/chat/service"
	"github.com/larek-tech/innohack/backend/internal/shared/database"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/rs/zerolog/log"
)

const (
	srvHeader   = "larek.tech"
	bodyLimitMb = 5
)

type Server struct {
	app *fiber.App
	cfg config.Config
}

func New(cfg config.Config) Server {
	if err := cfg.Server.Validate(); err != nil {
		panic(pkg.WrapErr(err, "config validation"))
	}

	app := fiber.New(fiber.Config{
		ServerHeader: srvHeader,
		BodyLimit:    bodyLimitMb * 1024 * 1024,
		ErrorHandler: errorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Server.GetOrigins(),
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Use(recovermw.New())

	app.Get("/", indexHanlder)
	app.Get("/health", healtCheckHandler(uuid.NewString()))
	app.Static("/static", "./static")

	s := Server{
		app: app,
		cfg: cfg,
	}

	pg := database.InitPostgres(context.Background(), cfg.Postgres.DSN)

	as := authService.New(pg, pg, &cfg.Auth.Oauth)
	cs := chatService.New()
	s.initModules(
		auth.New(as),
		chat.New(cs),
	)

	return s
}

func (s *Server) Serve() {
	go s.listenHttp(strconv.Itoa(s.cfg.Server.Port))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	if err := s.app.Shutdown(); err != nil {
		log.Err(pkg.WrapErr(err)).Msg("graceful shutdown")
	}
}

func (s *Server) listenHttp(port string) {
	addr := net.JoinHostPort("0.0.0.0", port)
	if err := s.app.Listen(addr); err != nil {
		log.Err(pkg.WrapErr(err)).Msg("application interrupted")
	}
}
