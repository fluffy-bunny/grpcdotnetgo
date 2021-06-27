package contextaccessor

import (
	"context"

	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type service struct {
	context context.Context
}

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("IContextAccessor")

// GetContextAccessorFromContainer from the Container
func GetContextAccessorFromContainer(ctn di.Container) IContextAccessor {
	return ctn.Get(diServiceName).(IContextAccessor)
}

// GetInternalGetContextAccessorFromContainer from the Container
func GetInternalGetContextAccessorFromContainer(ctn di.Container) IInternalContextAccessor {
	return ctn.Get(diServiceName).(IInternalContextAccessor)
}

// ContextAccessor adds service to the DI container
func AddContextAccessor(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceName).
		Msg("IoC: AddContextAccessor")

	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &service{}, nil
		},
	})
}
