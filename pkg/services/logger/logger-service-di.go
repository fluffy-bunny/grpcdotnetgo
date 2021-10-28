package logger

import (
	"reflect"

	contractsContextAccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/contextaccessor"
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// AddScopedILogger adds service to the DI container
func AddScopedILogger(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddScopedLogger")
	loggerContracts.AddScopedILoggerByFunc(builder, reflect.TypeOf(&loggerService{}),
		func(ctn di.Container) (interface{}, error) {
			contextAccessor := contractsContextAccessor.GetIContextAccessorFromContainer(ctn)
			logger := zerolog.Ctx(contextAccessor.GetContext())
			return &loggerService{
				Logger: logger,
			}, nil
		})
}

var diServiceNameILoggerSingleton = grpcdotnetgoutils.GenerateUnqueServiceName("ILogger-Singleton")

//GetSingletonLoggerFromContainer ...
func GetSingletonLoggerFromContainer(ctn di.Container) loggerContracts.ILogger {
	service := ctn.Get(diServiceNameILoggerSingleton).(loggerContracts.ILogger)
	return service
}

// AddSingletonILogger adds service to the DI container
func AddSingletonILogger(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceNameILoggerSingleton).
		Msg("IoC: AddSingletonLogger")
	builder.Add(di.Def{
		Name:  diServiceNameILoggerSingleton,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			logger := log.With().Logger()
			return &loggerService{
				Logger: &logger,
			}, nil
		},
	})
}
