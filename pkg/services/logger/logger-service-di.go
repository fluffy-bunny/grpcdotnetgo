package logger

import (
	"reflect"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// AddScopedILogger adds service to the DI container
func AddScopedILogger(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddScopedLogger")
	contracts_logger.AddScopedILoggerByFunc(builder, reflect.TypeOf(&loggerService{}),
		func(ctn di.Container) (interface{}, error) {
			request := contracts_request.GetIRequestFromContainer(ctn)
			ctx := request.GetContext()
			logger := zerolog.Ctx(ctx)
			return &loggerService{
				Logger: logger,
			}, nil
		})
}

var diServiceNameILoggerSingleton = grpcdotnetgoutils.GenerateUnqueServiceName("ILogger-Singleton")

// AddSingletonILogger adds service to the DI container
func AddSingletonILogger(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceNameILoggerSingleton).
		Msg("IoC: AddSingletonILoggerByFunc")

	contracts_logger.AddSingletonISingletonLoggerByFunc(builder, reflect.TypeOf(&loggerService{}),
		func(ctn di.Container) (interface{}, error) {
			logger := log.With().Logger()
			return &loggerService{
				Logger: &logger,
			}, nil
		})
}
