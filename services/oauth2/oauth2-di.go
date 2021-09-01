package oauth2

import (
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"

	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("IOauth2")

// GetOauth2ServiceFromContainer from the Container
func GetOauth2ServiceFromContainer(ctn di.Container) IOauth2 {
	service := ctn.Get(diServiceName).(IOauth2)
	return service
}

// AddServiceProvider adds service to the DI container
func AddOauth2Service(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceName).
		Msg("IoC: AddOauth2Service")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {

			return &service{
				ContextAccessor: contextaccessor.GetContextAccessorFromContainer(ctn),
				Logger:          servicesLogger.GetScopedLoggerFromContainer(ctn),
			}, nil
		},
	})
}
