// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"net"
	"reflect"

	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	exampleServices "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services"
	handlerGreeterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/helloworld/handler"
	singletonService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/singleton"
	transientService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/transient"
	dicontext_middleware "github.com/fluffy-bunny/grpcdotnetgo/middleware/dicontext"
	logger_middleware "github.com/fluffy-bunny/grpcdotnetgo/middleware/logger"
	grpcdotnetgreflect "github.com/fluffy-bunny/grpcdotnetgo/reflect"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	port = ":40051"
)

var version = "development"

func main() {
	fmt.Println("Version:\t", version)
	config := &internal.Config{}
	ReadViperConfig(internal.ConfigDefaultYaml, &config)

	// Create a Builder with the default scopes (App, Request, SubRequest).
	dotNetGoBuilder, err := grpcdotnetgo.NewDotNetGoBuilder()
	if err != nil {
		panic(err)
	}

	handlerGreeterService.AddGreeterService(dotNetGoBuilder.Builder)
	singletonService.AddSingletonService(dotNetGoBuilder.Builder)

	transientService.AddTransientService(dotNetGoBuilder.Builder)
	transientService.AddTransientService2(dotNetGoBuilder.Builder)

	dotNetGoBuilder.Build()

	ctn := grpcdotnetgo.GetContainer()
	inter := grpcdotnetgreflect.GetInterfaceReflectType((*exampleServices.ISomething)(nil))

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
}

func exampleAuthFunc(ctx context.Context) (context.Context, error) {

	return ctx, nil
}
