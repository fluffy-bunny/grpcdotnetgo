package oidc

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

func GetOIDCConfigAccessorFromContainer(ctn di.Container) IOIDCConfigAccessor {
	obj := ctn.GetByType(TypeIOIDCConfigAccessor).(IOIDCConfigAccessor)
	return obj
}

// AddIOIDCConfigAccessor adds service to the DI container
func AddIOIDCConfigAccessor(builder *di.Builder, obj interface{}) {
	log.Info().
		Msg("IoC: AddIOIDCConfigAccessor")
	types := di.NewTypeSet()
	types.Add(TypeIOIDCConfigAccessor)

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(obj),
		Build: func(ctn di.Container) (interface{}, error) {
			return obj, nil
		},
	})
}
