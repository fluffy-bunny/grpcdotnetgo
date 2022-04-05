package welcome

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// AddOneTimeWelcomeJobProvider adds service to the DI container
func AddOneTimeWelcomeJobProvider(builder *di.Builder) {
	backgroundtasksContracts.AddSingletonIJobsProvider(builder,
		reflect.TypeOf(&service{}))
}
