// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"

	grpcdotnetgocore "github.com/fluffy-bunny/grpcdotnetgo/core"
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	handlerGreeterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/helloworld/handler"
	singletonService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/singleton"
	transientService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/transient"
	dicontext_middleware "github.com/fluffy-bunny/grpcdotnetgo/middleware/dicontext"
	logger_middleware "github.com/fluffy-bunny/grpcdotnetgo/middleware/logger"
	runtime "github.com/fluffy-bunny/grpcdotnetgo/runtime"
	di "github.com/fluffy-bunny/sarulabsdi"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
)

var version = "development"

type Startup struct {
	port int
}

func (s *Startup) Startup() {
}
func (s *Startup) SetPort(port int) {
	s.port = port
}
func (s *Startup) GetPort() int {
	return s.port
}
func (s *Startup) ConfigureServices(builder *di.Builder) {
	handlerGreeterService.AddGreeterService(builder)
	singletonService.AddSingletonService(builder)

	transientService.AddTransientService(builder)
	transientService.AddTransientService2(builder)

}
func (s *Startup) Configure(
	container di.Container,
	unaryServerInterceptorBuilder *grpcdotnetgocore.UnaryServerInterceptorBuilder) {
	unaryServerInterceptorBuilder.Use(grpc_ctxtags.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(logger_middleware.EnsureContextLoggingUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(logger_middleware.EnsureCorrelationIDUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(dicontext_middleware.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(logger_middleware.LoggingUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(grpc_auth.UnaryServerInterceptor(exampleAuthFunc))
}
func (s *Startup) RegisterGRPCEndpoints(server *grpc.Server) {
	pb.RegisterGreeterServerDI(server)
}

func main() {
	runtime.SetVersion(version)
	fmt.Println("Version:\t", version)
	config := &internal.Config{}
	ReadViperConfig(internal.ConfigDefaultYaml, &config)

	runtime.Start(&Startup{})

}

func exampleAuthFunc(ctx context.Context) (context.Context, error) {

	return ctx, nil
}
