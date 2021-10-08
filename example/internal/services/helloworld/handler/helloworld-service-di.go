package handler

import (
	"reflect"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"

	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	servicesConfig "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/config"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddGreeterService adds service to the DI container
func AddGreeterService(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddGreeterService")
	types := di.NewTypeSet()
	types.Add(pb.TypeIGreeterService)

	builder.Add(di.Def{
		Scope:            di.Request,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service{}),
		Build: func(ctn di.Container) (interface{}, error) {
			return &Service{
				config:          servicesConfig.GetConfigFromContainer(ctn),
				ContextAccessor: contextaccessor.GetContextAccessorFromContainer(ctn),
				ClaimsPrincipal: claimsprincipalContracts.GetIClaimsPrincipalFromContainer(ctn),
				Logger:          servicesLogger.GetScopedLoggerFromContainer(ctn),
			}, nil
		},
	})
}

// AddGreeter2Service adds service to the DI container
func AddGreeter2Service(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddGreeter2Service")
	types := di.NewTypeSet()
	types.Add(pb.TypeIGreeter2Service)

	builder.Add(di.Def{
		Scope:            di.Request,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service2{}),
		Build: func(ctn di.Container) (interface{}, error) {
			return &Service2{
				config:          servicesConfig.GetConfigFromContainer(ctn),
				ContextAccessor: contextaccessor.GetContextAccessorFromContainer(ctn),
				ClaimsPrincipal: claimsprincipalContracts.GetIClaimsPrincipalFromContainer(ctn),
				Logger:          servicesLogger.GetScopedLoggerFromContainer(ctn),
			}, nil
		},
	})
}
