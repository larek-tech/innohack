package config

import (
	"github.com/ilyakaznacheev/cleanenv"

	auth "github.com/larek-tech/innohack/backend/internal/auth/config"
	server "github.com/larek-tech/innohack/backend/internal/server/config"
	"github.com/larek-tech/innohack/backend/internal/shared/database"
	"github.com/larek-tech/innohack/backend/pkg"
)

type Config struct {
	Server   *server.Config           `yaml:"server"`
	Postgres *database.PostgresConfig `yaml:"postgres"`
	Auth     *auth.Config       `yaml:"auth"`
}

func MustNewConfig(path string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(pkg.WrapErr(err, "load config"))
	}
	return cfg
}
