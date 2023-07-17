package config

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
)

// ExampleConfig type
type ExampleConfig struct {
	Port               int  `mapstructure:"PORT"`
	RESTPort           int  `mapstructure:"REST_PORT"`
	GRPCGatewayEnabled bool `json:"grpcGatewayEnabled" mapstructure:"GRPC_GATEWAY_ENABLED"`

	Mode             string          `mapstructure:"MODE"`
	OIDCConfig       oidc.OIDCConfig `mapstructure:"OIDC_CONFIG"`
	EnableTransient2 bool            `mapstructure:"ENABLE_TRANSIENT_2"`
}

// Config type
type Config struct {
	Environment string        `mapstructure:"APPLICATION_ENVIRONMENT"`
	Example     ExampleConfig `mapstructure:"EXAMPLE"`
}

// GetOIDCConfig gets config
func (c *Config) GetOIDCConfig() oidc.IOIDCConfig {
	return &c.Example.OIDCConfig
}

// ConfigDefaultJSON default yaml
var ConfigDefaultJSON = []byte(`
{
	"APPLICATION_ENVIRONMENT": "in-environment",
	"EXAMPLE": {
	  "ENABLE_TRANSIENT_2": true,
	  "PORT": 1111,
	  "GRPC_GATEWAY_ENABLED": false,
	  "REST_PORT": 1112,
	  "OIDC_CONFIG": {
		"AUTHORITY": "https://in-environment/",
		"CRON_REFRESH_SCHEDULE": "@every 0h1m0s"
	  }
	}
  }
`)
