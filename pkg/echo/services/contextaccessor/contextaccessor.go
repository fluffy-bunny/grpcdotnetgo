package contextaccessor

import (
	contracts_contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/contextaccessor"

	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type (
	service struct {
		context echo.Context
	}
)

func assertImplementation() {
	var _ contracts_contextaccessor.IInternalEchoContextAccessor = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddScopedIEchoContextAccessor registers the *service as a singleton.
func AddScopedIEchoContextAccessor(builder *di.Builder) {
	log.Info().Str("DI", "IInternalEchoContextAccessor - SCOPED").Send()
	contracts_contextaccessor.AddScopedIInternalEchoContextAccessor(builder, reflectType)
	log.Info().Str("DI", "IEchoContextAccessor - SCOPED").Send()
	contracts_contextaccessor.AddScopedIEchoContextAccessorByFunc(builder, reflectType, func(ctn di.Container) (interface{}, error) {
		internal := contracts_contextaccessor.GetIInternalEchoContextAccessorFromContainer(ctn)
		return internal.(contracts_contextaccessor.IEchoContextAccessor), nil
	})
}

func (s *service) SetContext(context echo.Context) {
	s.context = context
}
func (s *service) GetContext() echo.Context {
	return s.context
}
