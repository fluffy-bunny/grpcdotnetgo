package asynqengine

import (
	"reflect"

	contracts_asynqengine "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/asynqengine"

	"github.com/alicebob/miniredis/v2"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	service struct {
		miniR *miniredis.Miniredis
	}
)

func assertImplementation() {
	var _ contracts_asynqengine.IAsynqEngine = (*service)(nil)
}

// AddSingletonIModularAuthMiddleware ...
func AddSingletonIModularAuthMiddleware(builder *di.Builder) {
	contracts_asynqengine.AddSingletonIAsynqEngine(builder, reflect.TypeOf(&service{}))
}

func (s *service) Ctor() {
	var err error
	s.miniR, err = miniredis.Run()
	if err != nil {
		panic(err)
	}
}
