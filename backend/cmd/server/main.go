package main

import (
	"github.com/larek-tech/innohack/backend/config"
	"github.com/larek-tech/innohack/backend/internal/server"
)

func main() {
	cfg := config.MustNewConfig("./config/config.yaml")
	srv := server.New(cfg)
	srv.Serve()
}
