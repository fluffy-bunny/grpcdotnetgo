package core

import (
	"fmt"

	"net"

	"github.com/fluffy-bunny/grpcdotnetgo"
	"github.com/rs/zerolog/log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func Start(startup IStartup) {
	// Create a Builder with the default scopes (App, Request, SubRequest).
	dotNetGoBuilder, err := grpcdotnetgo.NewDotNetGoBuilder()
	if err != nil {
		panic(err)
	}
	startup.ConfigureServices(dotNetGoBuilder.Builder)
	dotNetGoBuilder.Build()
	unaryServerInterceptorBuilder := UnaryServerInterceptorBuilder{}
	startup.Configure(grpcdotnetgo.GetContainer(), &unaryServerInterceptorBuilder)

	port := fmt.Sprintf(":%v", startup.GetPort())
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Err(err).
			Str("port", port).
			Msg("failed to listen:")
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryServerInterceptorBuilder.UnaryServerInterceptors...,
		)),
	)
	startup.RegisterGRPCEndpoints(grpcServer)

	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve: ")
	}
}
