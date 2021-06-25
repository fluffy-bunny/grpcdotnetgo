package serviceprovider

import (
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	di "github.com/sarulabs/di/v2"
)

// Define an object in the App scope.

var diServiceName = "di-req-service-provider"

// GetDIServiceProviderFromContainer from the Container
func GetServiceProviderFromContainer(ctn di.Container) IServiceProvider {
	service := ctn.Get(diServiceName).(IServiceProvider)
	return service
}

// AddServiceProvider adds service to the DI container
func AddServiceProvider(builder *di.Builder) {
	log.Info().Msg("IoC: ServiceProvider")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			contextAccessor := contextaccessor.GetContextAccessorFromContainer(ctn)
			logger := zerolog.Ctx(contextAccessor.GetContext())
			return &serviceProvider{
				Logger:    logger,
				Container: ctn,
			}, nil
		},
	})
}
