package transient

import (
	"reflect"

	contracts_transient "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/transient"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

var (
	rtService  = reflect.TypeOf(&service{})
	rtService2 = reflect.TypeOf(&service2{})
)

// AddTransientITransient adds service to the DI container
func AddTransientITransient(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddTransientITransient")
	contracts_transient.AddTransientITransient(builder, rtService)
}

// AddTransientITransient2 adds service to the DI container
func AddTransientITransient2(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddTransientService2")
	contracts_transient.AddTransientITransient(builder, rtService2)
}
