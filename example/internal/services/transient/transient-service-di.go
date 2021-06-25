package handler

import (
	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	singletonServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/singleton-serviceprovider"
	"github.com/rs/zerolog/log"
	di "github.com/sarulabs/di/v2"
)

// Define an object in the App scope.
var diServiceName = "di-transient-service"

// GetTransientServiceFromContainer from the Container
func GetTransientService() *Service {
	return GetTransientServiceFromContainer(grpcdotnetgo.GetContainer())
}

// GetTransientServiceFromContainer from the Container
func GetTransientServiceFromContainer(ctn di.Container) *Service {
	return ctn.Get(diServiceName).(*Service)
}

// GreeterAddTransientServiceService adds service to the DI container
func AddTransientService(builder *di.Builder) {
	log.Info().Msg("IoC: AddTransientService")
	builder.Add(di.Def{
		Name:     diServiceName,
		Scope:    di.App,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &Service{
				ServiceProvider: singletonServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}, nil
		},
	})
}
