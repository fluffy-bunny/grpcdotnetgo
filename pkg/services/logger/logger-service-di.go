package logger

import (
	"reflect"

	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceNameILoggerScoped = grpcdotnetgoutils.GenerateUnqueServiceName("ILogger-Scoped")

var (
	reflectTypeILogger = di.GetInterfaceReflectType((*ILogger)(nil))
)

// GetScopedLoggerFromContainer from the Container
func GetScopedLoggerFromContainer(ctn di.Container) ILogger {
	service := ctn.GetByType(reflectTypeILogger).(ILogger)
	return service
}

// AddScopedLogger adds service to the DI container
func AddScopedLogger(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceNameILoggerScoped).
		Msg("IoC: AddScopedLogger")
	implementedTypes := di.NewTypeSet()
	implementedTypes.Add(reflectTypeILogger)
	builder.Add(di.Def{
		Name:             diServiceNameILoggerScoped,
		Type:             reflect.TypeOf(&loggerService{}),
		ImplementedTypes: implementedTypes,
		Scope:            di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			contextAccessor := contextaccessor.GetContextAccessorFromContainer(ctn)
			logger := zerolog.Ctx(contextAccessor.GetContext())
			return &loggerService{
				Logger: logger,
			}, nil
		},
	})
}

var diServiceNameILoggerSingleton = grpcdotnetgoutils.GenerateUnqueServiceName("ILogger-Singleton")

//GetSingletonLoggerFromContainer ...
func GetSingletonLoggerFromContainer(ctn di.Container) ILogger {
	service := ctn.Get(diServiceNameILoggerSingleton).(ILogger)
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
