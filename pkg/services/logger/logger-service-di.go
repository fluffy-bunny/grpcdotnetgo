package logger

import (
	"reflect"

	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceNameILoggerScoped = grpcdotnetgoutils.GenerateUnqueServiceName("ILogger-Scoped")

var (
	reflectTypeILogger = di.GetInterfaceReflectType((*loggerContracts.ILogger)(nil))
)

// AddScopedLogger adds service to the DI container
func AddScopedLogger(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceNameILoggerScoped).
		Msg("IoC: AddScopedLogger")

	loggerContracts.AddScopedILoggerByFunc(builder, reflect.TypeOf(&loggerService{}),
		func(ctn di.Container) (interface{}, error) {
			contextAccessor := contextaccessor.GetContextAccessorFromContainer(ctn)
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

// AddSingletonLogger adds service to the DI container
func AddSingletonLogger(builder *di.Builder) {
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
