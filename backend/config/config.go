package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/larek-tech/innohack/backend/internal/server"
)

type Config struct {
	Server *server.Config `yaml:"http"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
