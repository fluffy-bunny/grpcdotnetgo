package asynq

import (
	"math"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
)

type config struct {
	consumerServiceName string
	producerServiceName string
	analyticsRate       float64
}

// An Option customizes the config.
type Option func(cfg *config)

func newConfig(opts ...Option) *config {
	cfg := &config{
		consumerServiceName: "asynq",
		producerServiceName: "asynq",
		// analyticsRate: globalconfig.AnalyticsRate(),
		analyticsRate: math.NaN(),
	}
	if utils.BoolEnv("DD_TRACE_ASYNQ_ANALYTICS_ENABLED", false) {
		cfg.analyticsRate = 1.0
	}

	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithServiceName sets the config service name to serviceName.
func WithServiceName(serviceName string) Option {
	return func(cfg *config) {
		cfg.consumerServiceName = serviceName
		cfg.producerServiceName = serviceName
	}
}

// WithAnalytics enables Trace Analytics for all started spans.
func WithAnalytics(on bool) Option {
	return func(cfg *config) {
		if on {
			cfg.analyticsRate = 1.0
		} else {
			cfg.analyticsRate = math.NaN()
		}
	}
}

// WithAnalyticsRate sets the sampling rate for Trace Analytics events
// correlated to started spans.
func WithAnalyticsRate(rate float64) Option {
	return func(cfg *config) {
		if rate >= 0.0 && rate <= 1.0 {
			cfg.analyticsRate = rate
		} else {
			cfg.analyticsRate = math.NaN()
		}
	}
}
