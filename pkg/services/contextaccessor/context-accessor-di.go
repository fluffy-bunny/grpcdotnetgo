package contextaccessor

import (
	"context"
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type service struct {
	context context.Context
}

var (
	rtIContextAccessor         = di.GetInterfaceReflectType((*IContextAccessor)(nil))
	rtIInternalContextAccessor = di.GetInterfaceReflectType((*IInternalContextAccessor)(nil))
)

// GetContextAccessorFromContainer from the Container
func GetContextAccessorFromContainer(ctn di.Container) IContextAccessor {
	obj := ctn.GetByType(rtIContextAccessor).(IContextAccessor)
	return obj
}

// GetInternalGetContextAccessorFromContainer from the Container
func GetInternalGetContextAccessorFromContainer(ctn di.Container) IInternalContextAccessor {
	obj := ctn.GetByType(rtIInternalContextAccessor).(IInternalContextAccessor)
	return obj
}

// AddContextAccessor adds service to the DI container
func AddContextAccessor(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddContextAccessor")
	types := di.NewTypeSet()

	types.Add(rtIContextAccessor)
	types.Add(rtIInternalContextAccessor)

	builder.Add(di.Def{
		Scope:            di.Request,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&service{}),
		Build: func(ctn di.Container) (interface{}, error) {
			return &service{}, nil
		},
	})
}
