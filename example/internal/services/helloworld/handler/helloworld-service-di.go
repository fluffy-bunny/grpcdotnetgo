package handler

import (
	"reflect"

	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// AddScopedIGreeterService adds service to the DI container
func AddScopedIGreeterService(builder *di.Builder) {
	pb.AddScopedIGreeterService(builder, reflect.TypeOf(&Service{}))
}

// AddScopedIGreeter2Service adds service to the DI container
func AddScopedIGreeter2Service(builder *di.Builder) {
	pb.AddScopedIGreeter2Service(builder, reflect.TypeOf(&Service2{}))
}
