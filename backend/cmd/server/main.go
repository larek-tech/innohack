package main

import (
	"time"

	"github.com/larek-tech/innohack/backend/config"
	server "github.com/larek-tech/innohack/backend/internal/_server"
)

func main() {
	time.LoadLocation("Europe/Moscow")
	cfg := config.MustNewConfig("./config/config.yaml")
	srv := server.New(cfg)
	srv.Serve()
}
