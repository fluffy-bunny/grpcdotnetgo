package counter

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// AddCronCounterJobProvider adds service to the DI container
func AddCronCounterJobProvider(builder *di.Builder) {
	backgroundtasksContracts.AddSingletonIJobsProvider(builder,
		reflect.TypeOf(&service{}))
}
