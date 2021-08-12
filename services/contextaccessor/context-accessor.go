package contextaccessor

import (
	"context"

	"github.com/fluffy-bunny/grpcdotnetgo/middleware/dicontext"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type IContextAccessor interface {
	GetContext() context.Context
	GetContainer() di.Container
}
type IInternalContextAccessor interface {
	IContextAccessor
	SetContext(context.Context)
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
