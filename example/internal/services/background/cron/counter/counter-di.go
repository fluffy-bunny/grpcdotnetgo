package counter

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddCronCounterJobProvider adds service to the DI container
func AddCronCounterJobProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddCronCounterJobProvider")
	backgroundtasksContracts.AddSingletonIJobsProviderByFunc(builder,
		reflect.TypeOf(&service{}), func(ctn di.Container) (interface{}, error) {
			obj := &service{
				Logger: servicesLogger.GetSingletonLoggerFromContainer(ctn),
			}

			return obj, nil
		})
}
