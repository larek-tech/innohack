package main

import (
	"github.com/larek-tech/innohack/backend/config"
	"github.com/larek-tech/innohack/backend/internal/server"
)

func main() {
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	srv, err := server.New(cfg.Server)
	if err != nil {
		panic(err)
	}
	err = srv.Serve()
	if err != nil {
		panic(err)
	}
}
