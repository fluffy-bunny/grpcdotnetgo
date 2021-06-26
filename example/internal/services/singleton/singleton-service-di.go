package singleton

import (
	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	singletonServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/singleton-serviceprovider"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("di-singleton-service")

// GetSingletonServiceFromContainer from the Container
func GetSingletonService() *Service {
	return GetSingletonServiceFromContainer(grpcdotnetgo.GetContainer())
}

// GetSingletonServiceFromContainer from the Container
func GetSingletonServiceFromContainer(ctn di.Container) *Service {
	return ctn.Get(diServiceName).(*Service)
}

// GreeterAddSingletonServiceService adds service to the DI container
func AddSingletonService(builder *di.Builder) {
	log.Info().Msg("IoC: AddSingletonService")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &Service{
				ServiceProvider: singletonServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}, nil
		},
	})
}
