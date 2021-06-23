package handler

import (
	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	singletonServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/singleton-serviceprovider"
	"github.com/rs/zerolog/log"
	"github.com/sarulabs/di"
)

// Define an object in the App scope.
var diServiceName = "di-singleton-service"

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
