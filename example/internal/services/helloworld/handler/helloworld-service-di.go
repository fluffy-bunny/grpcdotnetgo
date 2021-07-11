package handler

import (
	"reflect"

	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	"github.com/fluffy-bunny/grpcdotnetgo/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// GreeterService adds service to the DI container
func AddGreeterService(builder *di.Builder) {
	log.Info().
		Msg("IoC: GreeterService")
	types := di.NewTypeSet()
	types.Add(pb.TypeIGreeterService)

	builder.Add(di.Def{
		Scope:            di.Request,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service{}),
		Build: func(ctn di.Container) (interface{}, error) {
			return &Service{
				ContextAccessor: contextaccessor.GetContextAccessorFromContainer(ctn),
				ClaimsPrincipal: claimsprincipal.GetClaimsPrincipalFromContainer(ctn),
				Logger:          servicesLogger.GetScopedLoggerFromContainer(ctn),
				ServiceProvider: servicesServiceProvider.GetScopedServiceProviderFromContainer(ctn),
			}, nil
		},
	})
}

// GreeterService adds service to the DI container
func AddGreeter2Service(builder *di.Builder) {
	log.Info().
		Msg("IoC: GreeterService")
	types := di.NewTypeSet()
	types.Add(pb.TypeIGreeter2Service)

	builder.Add(di.Def{
		Scope:            di.Request,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service{}),
		Build: func(ctn di.Container) (interface{}, error) {
			return &Service{
				ContextAccessor: contextaccessor.GetContextAccessorFromContainer(ctn),
				ClaimsPrincipal: claimsprincipal.GetClaimsPrincipalFromContainer(ctn),
				Logger:          servicesLogger.GetScopedLoggerFromContainer(ctn),
				ServiceProvider: servicesServiceProvider.GetScopedServiceProviderFromContainer(ctn),
			}, nil
		},
	})
}
