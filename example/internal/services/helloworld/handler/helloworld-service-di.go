package handler

import (
	"reflect"

	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// AddGreeterEndpointRegistration adds service to the DI container
func AddGreeterEndpointRegistration(builder *di.Builder) {
	pb.AddGreeterEndpointRegistrationV2(builder, reflect.TypeOf(&Service{}))
	//pb.AddScopedIGreeterService(builder, reflect.TypeOf(&Service{}))
}

// AddGreeter2EndpointRegistration adds service to the DI container
func AddGreeter2EndpointRegistration(builder *di.Builder) {
	pb.AddGreeter2EndpointRegistrationV2(builder, reflect.TypeOf(&Service2{}))
	//pb.AddScopedIGreeter2Service(builder, reflect.TypeOf(&Service2{}))
}
