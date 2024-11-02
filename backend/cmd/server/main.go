package main

import (
	"time"

	"github.com/larek-tech/innohack/backend/config"
	"github.com/larek-tech/innohack/backend/internal/server"
)

func main() {
	time.LoadLocation("Europe/Moscow")
	cfg := config.MustNewConfig("./config/config.yaml")
	srv := server.New(cfg)
	srv.Serve()
}
