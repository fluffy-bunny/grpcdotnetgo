package singleton

import (
	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("di-singleton-service")

// GetSingletonServiceFromContainer from the Container
func GetSingletonService() *service {
	return GetSingletonServiceFromContainer(grpcdotnetgo.GetContainer())
}

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
				ServiceProvider: servicesServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}, nil
		},
	})
}
