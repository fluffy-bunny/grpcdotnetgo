package welcome

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddOneTimeWelcomeJobProvider adds service to the DI container
func AddOneTimeWelcomeJobProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddOneTimeWelcomeJobProvider")

	backgroundtasksContracts.AddSingletonIJobsProvider(builder,
		reflect.TypeOf(&service{}))
}
