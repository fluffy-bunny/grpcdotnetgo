package contextaccessor

import (
	"context"
	"reflect"

	contractsContextAccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/contextaccessor"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type service struct {
	context context.Context
}

// AddScopedContextAccessor adds service to the DI container
func AddScopedContextAccessor(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddContextAccessor")
	di.AddScopedWithImplementedTypes(builder, reflect.TypeOf(&service{}),
		contractsContextAccessor.ReflectTypeIContextAccessor,
		contractsContextAccessor.ReflectTypeIInternalContextAccessor)
}

func (s *service) GetContainer() di.Container {
	return dicontext.GetRequestContainer(s.context)
}
func (s *service) GetContext() context.Context {
	return s.context
}
func (s *service) SetContext(ctx context.Context) {
	s.context = ctx
}
