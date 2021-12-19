package counter

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddCronCounterJobProvider adds service to the DI container
func AddCronCounterJobProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddCronCounterJobProvider")
	backgroundtasksContracts.AddSingletonIJobsProvider(builder,
		reflect.TypeOf(&service{}))
}
