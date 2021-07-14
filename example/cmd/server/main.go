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
	grpc_recovery "github.com/fluffy-bunny/grpcdotnetgo/middleware/recovery"
	grpcDIProtoError "github.com/fluffy-bunny/grpcdotnetgo/proto/error"
	runtime "github.com/fluffy-bunny/grpcdotnetgo/runtime"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gogo/protobuf/gogoproto"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	handlerGreeterService.AddGreeter2Service(builder)

	singletonService.AddSingletonService(builder)

	transientService.AddTransientService(builder)
	transientService.AddTransientService2(builder)

}
func (s *Startup) Configure(
	container di.Container,
	unaryServerInterceptorBuilder *grpcdotnetgocore.UnaryServerInterceptorBuilder) {

	//var recoveryFunc grpc_recovery.RecoveryHandlerFunc
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerUnary(recoveryUnaryFunc),
	}
	unaryServerInterceptorBuilder.Use(grpc_ctxtags.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(logger_middleware.EnsureContextLoggingUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(logger_middleware.EnsureCorrelationIDUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(dicontext_middleware.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(logger_middleware.LoggingUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(grpc_auth.UnaryServerInterceptor(exampleAuthFunc))
	unaryServerInterceptorBuilder.Use(grpc_recovery.UnaryServerInterceptor(recoveryOpts...))

}
func (s *Startup) RegisterGRPCEndpoints(server *grpc.Server) {
	pb.RegisterGreeterServerDI(server)
	pb.RegisterGreeter2ServerDI(server)
}

func main() {
	d := gogoproto.E_GoprotoEnumStringer
	if d == nil {
		panic("boo hoo")
	}
	runtime.SetVersion(version)
	fmt.Println("Version:\t", version)
	config := &internal.Config{}
	ReadViperConfig(internal.ConfigDefaultYaml, &config)

	runtime.Start(&Startup{})

}

func exampleAuthFunc(ctx context.Context) (context.Context, error) {

	return ctx, nil
}
func recoveryUnaryFunc(fullMethodName string, p interface{}) (interface{}, error) {
	fmt.Printf("p: %+v\n", p)

	replyFunc, ok := pb.M_helloworldFullMethodNameWithErrorResponseMap[fullMethodName]
	if ok {
		reply, ok2 := replyFunc().(grpcDIProtoError.IError)
		if ok2 {
			myError := reply.GetError()
			myError.Code = 503
			myError.Message = "Unexpected error2"
			return reply, nil
		}
		ok = false

	}
	if !ok {
		return nil, status.Errorf(codes.Internal, "Unexpected error1")
	}
	return nil, nil
}
