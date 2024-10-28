package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/larek-tech/innohack/backend/internal/server/config"
	"github.com/larek-tech/innohack/backend/pkg"
)

type Config struct {
	Server *server.Config `yaml:"server"`
}

func MustNewConfig(path string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(pkg.WrapErr(err, "load config"))
	}
	return cfg
}
