package internal

import (
	"github.com/fluffy-bunny/grpcdotnetgo/middleware/oidc"
)

type Config struct {
	Environment string `mapstructure:"APPLICATION_ENVIRONMENT"`
	GRPCPort    int    `mapstructure:"GRPC_PORT"`
	//	Environment string `mapstructure:"APPLICATION_ENVIRONMENT"`
	Mode             string          `mapstructure:"MODE"`
	OIDCConfig       oidc.OIDCConfig `mapstructure:"OIDC_CONFIG"`
	EnableTransient2 bool            `mapstructure:"ENABLE_TRANSIENT_2"`
}

func (c *Config) GetOIDCConfig() oidc.IOIDCConfig {
	return &c.OIDCConfig
}

// ConfigDefaultYaml default yaml
var ConfigDefaultYaml = []byte(`
APPLICATION_ENVIRONMENT: in-environment
GRPC_PORT: 1111
ENABLE_TRANSIENT_2: true
OIDC_CONFIG:
  Authority: "https://in-environment/"
  CRON_REFRESH_SCHEDULE: "@every 0h1m0s"
  B: dd
`)
