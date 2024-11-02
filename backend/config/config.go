package config

import (
	"github.com/ilyakaznacheev/cleanenv"

	server "github.com/larek-tech/innohack/backend/internal/server/config"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/storage/postgres"
	"github.com/larek-tech/innohack/backend/pkg/tracing"
)

type Config struct {
	Server   *server.Config   `yaml:"server"`
	Postgres *postgres.Config `yaml:"postgres"`
	Jaeger   *tracing.Config  `yaml:"jaeger"`
}

func MustNewConfig(path string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(pkg.WrapErr(err, "load config"))
	}
	return cfg
}
