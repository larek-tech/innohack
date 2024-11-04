package main

import (
	"time"

	"github.com/larek-tech/innohack/backend/config"
	_ "github.com/larek-tech/innohack/backend/docs"
	server "github.com/larek-tech/innohack/backend/internal/_server"
)

// @title			MTS AI Docs
// @version		1.0
// @description	Документация для сервиса решения команды MISIS Banach Space к задаче MTS AI Docs.
// @host			localhost:9999
// @BasePath		/
func main() {
	time.LoadLocation("Europe/Moscow")
	cfg := config.MustNewConfig("./config/config.yaml")
	srv := server.New(cfg)
	srv.Serve()
}
