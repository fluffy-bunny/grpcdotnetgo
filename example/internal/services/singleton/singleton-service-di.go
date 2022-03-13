package singleton

import (
	"reflect"

	contracts_singleton "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/singleton"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

var (
	rtGetType = reflect.TypeOf(&service{})
)

// AddSingletonISingleton adds service to the DI container
func AddSingletonISingleton(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddSingletonService")
	contracts_singleton.AddSingletonISingleton(builder, rtGetType)
}
