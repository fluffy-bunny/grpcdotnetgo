package serviceprovider

import (
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceNameIServiceProviderScoped = grpcdotnetgoutils.GenerateUnqueServiceName("IServiceProvider-Scoped")

// GetScopedServiceProviderFromContainer from the Container
func GetScopedServiceProviderFromContainer(ctn di.Container) IServiceProvider {
	service := ctn.Get(diServiceNameIServiceProviderScoped).(IServiceProvider)
	return service
}

// AddServiceProvider adds service to the DI container
func AddScopedServiceProvider(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceNameIServiceProviderScoped).
		Msg("IoC: AddScopedServiceProvider")
	builder.Add(di.Def{
		Name:  diServiceNameIServiceProviderScoped,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &serviceProvider{
				Logger:    servicesLogger.GetScopedLoggerFromContainer(ctn),
				Container: ctn,
			}, nil
		},
	})
}

// Define an object in the App scope.
var diServiceNameIServiceProviderSingleton = grpcdotnetgoutils.GenerateUnqueServiceName("IServiceProvider-Singleton")

// GetSingletonServiceProviderFromContainer from the Container
func GetSingletonServiceProviderFromContainer(ctn di.Container) IServiceProvider {
	service := ctn.Get(diServiceNameIServiceProviderSingleton).(IServiceProvider)
	return service
}

// AddSingletonServiceProvider adds service to the DI container
func AddSingletonServiceProvider(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceNameIServiceProviderSingleton).
		Msg("IoC: AddSingletonServiceProvider")
	builder.Add(di.Def{
		Name:  diServiceNameIServiceProviderSingleton,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &serviceProvider{
				Logger:    &zerolog.Logger{},
				Container: ctn,
			}, nil
		},
	})
}
