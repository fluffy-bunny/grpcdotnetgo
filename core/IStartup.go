package core

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

type ICoreConfig interface {
	GetPort() int
}
type IStartup interface {
	ConfigureServices(builder *di.Builder)
	Configure(
		container di.Container,
		unaryServerInterceptorBuilder *UnaryServerInterceptorBuilder)
	GetPort() int
	RegisterGRPCEndpoints(server *grpc.Server)
}

type IUnaryServerInterceptorBuilder interface {
	Use(intercepter grpc.UnaryServerInterceptor)
}
