package server

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrInvalidPort    = errors.New("invalid http server port, must be in range 1000 - 65000")
	ErrMissingOrigins = errors.New("no origins specified")
)

type Config struct {
	Port         int      `yaml:"port"`
	AllowOrigins []string `yaml:"allow_origins"`
	JwtSecret    string   `yaml:"jwt_secret"`
}

func (c *Config) Validate() error {
	if c.Port < 1000 || c.Port > 65000 {
		return ErrInvalidPort
	}

	if len(c.AllowOrigins) == 0 {
		return ErrMissingOrigins
	}

	for _, origin := range c.AllowOrigins {
		_, err := url.ParseRequestURI(origin)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) GetOrigins() string {
	return strings.Join(c.AllowOrigins, ",")
}
