package myextension

import (
	"errors"

	"go.opentelemetry.io/collector/component"
)

var _ component.Config = (*Config)(nil)

type Config struct {
	URL string `mapstructure:"url"`
}

func (c *Config) Validate() error {
	if c.URL == "" {
		return errors.New("url is required")
	}
	return nil
}
