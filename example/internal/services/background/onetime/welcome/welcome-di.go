package welcome

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddOneTimeWelcomeJobProvider adds service to the DI container
func AddOneTimeWelcomeJobProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddOneTimeWelcomeJobProvider")

	backgroundtasksContracts.AddSingletonIJobsProviderByFunc(builder,
		reflect.TypeOf(&service{}), func(ctn di.Container) (interface{}, error) {
			obj := &service{
				Logger: servicesLogger.GetSingletonLoggerFromContainer(ctn),
			}
			return obj, nil
		})
}
