package contextaccessor

import (
	"context"

	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	"github.com/rs/zerolog/log"
	di "github.com/sarulabs/di/v2"
)

type service struct {
	context context.Context
}

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("context-accessor")

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
	log.Info().Msg("IoC: ContextAccessor")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &service{}, nil
		},
	})
}
