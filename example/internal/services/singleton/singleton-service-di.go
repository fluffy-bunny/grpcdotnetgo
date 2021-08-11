package singleton

import (
	servicesConfig "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/config"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.

var diServiceName = di.GenerateUniqueServiceKeyFromInterface(&service{})

// GetSingletonServiceFromContainer from the Container
func GetSingletonServiceFromContainer(ctn di.Container) *service {
	return ctn.Get(diServiceName).(*service)
}

// GreeterAddSingletonServiceService adds service to the DI container
func AddSingletonService(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceName).
		Msg("IoC: AddSingletonService")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &service{
				config: servicesConfig.GetConfigFromContainer(ctn),
			}, nil
		},
	})
}
