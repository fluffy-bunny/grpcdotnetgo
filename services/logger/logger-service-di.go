package logger

import (
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("ILogger")

// GetDIRequestLoggerFromContainer from the Container
func GetRequestLoggerFromContainer(ctn di.Container) ILogger {
	service := ctn.Get(diServiceName).(ILogger)
	return service
}

// AddRequestLogger adds service to the DI container
func AddRequestLogger(builder *di.Builder) {
	log.Info().Msg("IoC: DIRequestLogger")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			contextAccessor := contextaccessor.GetContextAccessorFromContainer(ctn)
			logger := zerolog.Ctx(contextAccessor.GetContext())
			return &LoggerService{
				Logger: logger,
			}, nil
		},
	})
}
