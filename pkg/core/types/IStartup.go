package types

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

type ICoreConfig interface {
	GetPort() int
}

type IStartup interface {
	GetConfigOptions() *ConfigOptions
	ConfigureServices(builder *di.Builder)
	Configure(unaryServerInterceptorBuilder IUnaryServerInterceptorBuilder)
	GetPort() int

	RegisterGRPCEndpoints(server *grpc.Server) []interface{}
	SetRootContainer(container di.Container)
}

type IUnaryServerInterceptorBuilder interface {
	GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor
	Use(intercepter grpc.UnaryServerInterceptor)
}
