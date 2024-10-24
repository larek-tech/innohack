package server

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"github.com/larek-tech/innohack/backend/templ/pages"
	"github.com/rs/zerolog/log"
)

type Server struct {
	router *fiber.App
	cfg    *Config
}

func New(cfg *Config, modules ...Module) (*Server, error) {
	nodeID := uuid.NewString()
	if cfg == nil {
		return nil, errors.New("invalid server config")
	}
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "larek",
		BodyLimit:    5 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.GetOriginsString(),
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"page": "index"})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"hostname": nodeID})
	})
	app.Static("/static", "./static")

	// not found route
	app.Use(func(c *fiber.Ctx) error {
		return adaptor.HTTPHandler(templ.Handler(pages.NotFound(c.Path())))(c)
	})

	fmt.Print(len(modules))
	for _, m := range modules {
		err := m.InitRoutes(app)
		log.Info().Msgf("Module %s initialized", m.Name())
		if err != nil {
			log.Err(err).Str("module", m.Name()).Msg("unable to init")
		}
	}

	return &Server{
		router: app,
		cfg:    cfg,
	}, nil
}

func (s *Server) Serve() error {
	if s.cfg == nil {
		return errors.New("unable to load config for server")
	}
	if s.router == nil {
		return errors.New("fiber app was not configured")
	}
	addr := "0.0.0.0:" + strconv.FormatInt(int64(s.cfg.Port), 10)

	return s.router.Listen(addr)
}
