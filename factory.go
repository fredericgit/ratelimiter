package ratelimiter

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/processor"
)

// NewFactory creates a factory for ratelimiter processor.
func NewFactory() processor.Factory {
		return processor.NewFactory(
		"ratelimiter",
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, component.StabilityLevelDevelopment))
}

func createDefaultConfig() component.Config {
	return &Config{
		LimitRate: "10000/1s",
		LimitBy: [],
	}
}

// createTracesProcessor creates a trace processor based on this config.
func createTracesProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	return newTracesProcessor(ctx, set, cfg.(*Config), nextConsumer)
}