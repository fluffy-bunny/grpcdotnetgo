package contextaccessor

import "context"

type IContextAccessor interface {
	GetContext() context.Context
}
type IInternalContextAccessor interface {
	IContextAccessor
	SetContext(context.Context)
}

func (s *service) GetContext() context.Context {
	return s.context
}
func (s *service) SetContext(ctx context.Context) {
	s.context = ctx
}
