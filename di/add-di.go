package di

import (
	"fmt"
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddByType adds service to the DI container
func AddByType(builder *di.Builder,
	scope string,
	build func(ctn di.Container) (interface{}, error),
	close func(obj interface{}) error,
	rootType reflect.Type,
	implementedTypes ...reflect.Type) {
	lv := log.Info().
		Str("rootType", fmt.Sprintf("%v", rootType))
	for _, rt := range implementedTypes {
		lv = lv.Str("implementedType", fmt.Sprintf("%v", rt))
	}
	lv.Msg("IoC: AddByType")

	types := di.NewTypeSet()
	for _, rt := range implementedTypes {
		types.Add(rt)
	}
	var unshared bool = false
	if scope == "transient" {
		scope = di.App
		unshared = true
	}
	def := di.Def{
		Scope:            scope,
		Unshared:         unshared,
		ImplementedTypes: types,
		Type:             rootType,
		Build:            build,
	}
	if close != nil {
		def.Close = close
	}

	builder.Add(def)
}

// AddSingletonByType adds service to the DI container
func AddSingletonByType(builder *di.Builder,
	build func(ctn di.Container) (interface{}, error),
	close func(obj interface{}) error,
	rootType reflect.Type,
	implementedTypes ...reflect.Type) {
	AddByType(builder, di.App, build, close, rootType, implementedTypes...)
}

// AddSingletonByType adds service to the DI container
func AddScopedByType(builder *di.Builder,
	build func(ctn di.Container) (interface{}, error),
	close func(obj interface{}) error,
	rootType reflect.Type,
	implementedTypes ...reflect.Type) {
	AddByType(builder, di.Request, build, close, rootType, implementedTypes...)
}

// AddTransientByType adds service to the DI container
func AddTransientByType(builder *di.Builder,
	build func(ctn di.Container) (interface{}, error),
	close func(obj interface{}) error,
	rootType reflect.Type,
	implementedTypes ...reflect.Type) {
	AddByType(builder, "transient", build, close, rootType, implementedTypes...)
}
