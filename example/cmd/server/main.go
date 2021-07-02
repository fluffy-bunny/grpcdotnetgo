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
	di "github.com/fluffy-bunny/sarulabsdi"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
)

const (
	port = 40051
)

var version = "development"

type Startup struct {
	port int
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
	fmt.Println("Version:\t", version)
	config := &internal.Config{}
	ReadViperConfig(internal.ConfigDefaultYaml, &config)

	grpcdotnetgocore.Start(&Startup{
		port: port,
	})
	/*
		ctn := grpcdotnetgo.GetContainer()
		inter := di.GetInterfaceReflectType((*exampleServices.ISomething)(nil))

		dd := ctn.GetManyByType(inter)
		for _, d := range dd {
			ds := d.(exampleServices.ISomething)
			ds.SetName("rabbit")
			log.Info().Msg(ds.GetName())
		}

		inter = reflect.TypeOf(&transientService.Service{}).Elem()
		dd = ctn.GetManyByType(inter)
		for _, d := range dd {
			ds := d.(*transientService.Service)
			ds.SetName("Cougar")
			log.Info().Msg(ds.GetName())
		}

		ss := singletonService.GetSingletonService()
		ss.SetName("test")
		log.Info().Msg(ss.GetName())

		ss2 := singletonService.GetSingletonService()
		log.Info().Msg(ss2.GetName())

		ts := transientService.GetTransientService()
		ts.SetName("test")
		log.Info().Msg(ts.GetName())

		ts2 := transientService.GetTransientService()
		log.Info().Msg(ts2.GetName())

		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal().Err(err).
				Str("port", port).
				Msg("failed to listen:")
		}
		grpcServer := grpc.NewServer(

			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				logger_middleware.EnsureContextLoggingUnaryServerInterceptor(),
				logger_middleware.EnsureCorrelationIDUnaryServerInterceptor(),
				dicontext_middleware.UnaryServerInterceptor(),
				logger_middleware.LoggingUnaryServerInterceptor(),
				grpc_auth.UnaryServerInterceptor(exampleAuthFunc),
			)),
		)
		pb.RegisterGreeterServerDI(grpcServer)

		log.Info().Msgf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("failed to serve: ")
		}
	*/
}

func exampleAuthFunc(ctx context.Context) (context.Context, error) {

	return ctx, nil
}
