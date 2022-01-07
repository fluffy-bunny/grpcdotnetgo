package handler

import (
	"reflect"

	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddScopedIGreeterService adds service to the DI container
func AddScopedIGreeterService(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddScopedIGreeterService")
	pb.AddScopedIGreeterService(builder, reflect.TypeOf(&Service{}))
}

// AddScopedIGreeter2Service adds service to the DI container
func AddScopedIGreeter2Service(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddScopedIGreeter2Service")
	pb.AddScopedIGreeter2Service(builder, reflect.TypeOf(&Service2{}))
}
