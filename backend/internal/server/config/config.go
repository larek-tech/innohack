package server

import (
	"errors"
	"net/url"
	"strings"

	"github.com/larek-tech/innohack/backend/pkg"
)

var (
	ErrInvalidPort    = errors.New("invalid http server port, must be in range 1000 - 65000")
	ErrMissingOrigins = errors.New("no origins specified")
)

type Config struct {
	Port         int      `yaml:"port"`
	AllowOrigins []string `yaml:"allow_origins"`
}

func (c *Config) Validate() error {
	if c.Port < 1000 || c.Port > 65000 {
		return pkg.WrapErr(ErrInvalidPort)
	}

	if len(c.AllowOrigins) == 0 {
		return pkg.WrapErr(ErrMissingOrigins)
	}

	for _, origin := range c.AllowOrigins {
		_, err := url.ParseRequestURI(origin)
		if err != nil {
			return pkg.WrapErr(err)
		}
	}
	return nil
}

func (c *Config) GetOrigins() string {
	return strings.Join(c.AllowOrigins, ",")
}
