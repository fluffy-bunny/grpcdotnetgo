package config

import (
	contracts_config "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/config"
	servicesGrpcDotNetGoConfig "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/config"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// GetConfigFromContainer from the Container
func GetConfigFromContainer(ctn di.Container) *contracts_config.Config {
	obj := servicesGrpcDotNetGoConfig.GetConfigFromContainer(ctn).(*contracts_config.Config)
	return obj
}
