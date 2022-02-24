package logger

import (
	"reflect"

	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
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

// AddSingletonILogger adds service to the DI container
func AddSingletonILogger(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddSingletonILogger")
	bldFunc := func(ctn di.Container) (interface{}, error) {
		logger := log.With().Logger()
		return &loggerService{
			Logger: &logger,
		}, nil
	}
	contracts_logger.AddSingletonISingletonLoggerByFunc(builder, reflect.TypeOf(&loggerService{}), bldFunc)
	contracts_logger.AddSingletonILoggerByFunc(builder, reflect.TypeOf(&loggerService{}), bldFunc)
}
