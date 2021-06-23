// Package main implements a server for Greeter service.
package main

import (
	"net"

	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	handlerGreeterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/helloworld/handler"
	dicontext_middleware "github.com/fluffy-bunny/grpcdotnetgo/middleware/dicontext"
	logger_middleware "github.com/fluffy-bunny/grpcdotnetgo/middleware/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	port = ":40051"
)

func main() {
	config := &internal.Config{}
	ReadViperConfig(internal.ConfigDefaultYaml, &config)

	// Create a Builder with the default scopes (App, Request, SubRequest).
	dotNetGoBuilder, err := grpcdotnetgo.NewDotNetGoBuilder()
	if err != nil {
		panic(err)
	}

	handlerGreeterService.AddGreeterService(dotNetGoBuilder.Builder)

	dotNetGoBuilder.Build()

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
		)),
	)
	pb.RegisterGreeterServerDI(grpcServer)

	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve: ")
	}
}
