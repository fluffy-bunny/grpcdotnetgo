package asynq

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyticsSettings(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		cfg := newConfig()
		assert.True(t, math.IsNaN(cfg.analyticsRate))
	})

	t.Run("enabled", func(t *testing.T) {
		cfg := newConfig(WithAnalytics(true))
		assert.Equal(t, 1.0, cfg.analyticsRate)
	})
}
