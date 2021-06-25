package serviceprovider

import (
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	di "github.com/sarulabs/di/v2"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("di-singleton-service-provider")

// GetSingletonServiceProviderFromContainer from the Container
func GetSingletonServiceProviderFromContainer(ctn di.Container) IServiceProvider {
	service := ctn.Get(diServiceName).(IServiceProvider)
	return service
}

// AddSingletonServiceProvider adds service to the DI container
func AddSingletonServiceProvider(builder *di.Builder) {
	log.Info().Msg("IoC: AddSingletonServiceProvider")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &serviceProvider{
				Logger:    &zerolog.Logger{},
				Container: ctn,
			}, nil
		},
	})
}
