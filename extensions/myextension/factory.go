package myextension

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func NewFactory() extension.Factory {
	return extension.NewFactory(
		component.MustNewType("myextension"),
		createDefaultConfig,
		create,
		component.StabilityLevelDevelopment,
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		URL: "https://example.com",
	}
}

func create(ctx context.Context, set extension.Settings, cfg component.Config) (extension.Extension, error) {
	return NewExtension(cfg.(*Config), set.Logger), nil
}
