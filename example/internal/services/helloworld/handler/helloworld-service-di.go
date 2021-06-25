package handler

import (
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	"github.com/fluffy-bunny/grpcdotnetgo/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	"github.com/rs/zerolog/log"
	di "github.com/sarulabs/di/v2"
)

// GreeterService adds service to the DI container
func AddGreeterService(builder *di.Builder) {
	log.Info().Msg("IoC: GreeterService")
	builder.Add(di.Def{
		Name:  pb.GetGreeterServiceName(),
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &Service{
				ContextAccessor: contextaccessor.GetContextAccessorFromContainer(ctn),
				ClaimsPrincipal: claimsprincipal.GetClaimsPrincipalFromContainer(ctn),
				Logger:          servicesLogger.GetRequestLoggerFromContainer(ctn),
				ServiceProvider: servicesServiceProvider.GetServiceProviderFromContainer(ctn),
			}, nil
		},
	})
}
