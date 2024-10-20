package server

import (
	"errors"
	"net/url"
	"strings"

	"go.uber.org/multierr"
)

var (
	ErrInvalidPort = errors.New("invalid http server port must be in range 1000 - 65000")
)

type Config struct {
	Domain       string   `yaml:"domain"`
	Port         int      `yaml:"port"`
	AllowOrigins []string `yaml:"allow_origins"`
}

func (c *Config) Validate() error {
	var (
		ac error = nil
	)
	if c.Port < 1000 || c.Port > 65000 {
		ac = multierr.Append(ac, ErrInvalidPort)
	}

	for _, origin := range c.AllowOrigins {
		_, err := url.ParseRequestURI(origin)
		if err != nil {
			ac = multierr.Append(ac, err)
		}
	}
	return ac
}

func (c *Config) GetOriginsString() string {
	if len(c.AllowOrigins) == 0 {
		return "*"
	}
	return strings.Join(c.AllowOrigins, ",")
}
