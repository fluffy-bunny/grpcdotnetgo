package scoped

import (
	"reflect"

	contracts_scoped "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/scoped"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

var (
	rtGetType = reflect.TypeOf(&service{})
)

// AddScopedIScoped adds service to the DI container
func AddScopedIScoped(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddScopedIScoped")
	contracts_scoped.AddScopedIScoped(builder, rtGetType)
}
