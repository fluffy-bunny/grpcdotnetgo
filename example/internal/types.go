package internal

import (
	grpcdotnetgocore "github.com/fluffy-bunny/grpcdotnetgo/core"
)

type Config struct {
	grpcdotnetgocore.Config `mapstructure:"CORE"`
	//	Environment string `mapstructure:"APPLICATION_ENVIRONMENT"`
	Mode string `mapstructure:"MODE"`
}

// ConfigDefaultYaml default yaml
var ConfigDefaultYaml = []byte(`
APPLICATION_ENVIRONMENT: in-environment
CORE:
  APPLICATION_ENVIRONMENT: in-environment
  GRPC_PORT: 1111
`)
