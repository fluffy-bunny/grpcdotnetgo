package types

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// ICoreConfig ...
type ICoreConfig interface {
	GetPort() int
}

// IStartup contract that matches asp.net core closely
type IStartup interface {
	GetConfigOptions() *ConfigOptions
	ConfigureServices(builder *di.Builder)
	Configure(unaryServerInterceptorBuilder IUnaryServerInterceptorBuilder)
	GetPort() int

	RegisterGRPCEndpoints(server *grpc.Server) []interface{}
	SetRootContainer(container di.Container)
}

// IUnaryServerInterceptorBuilder ...
type IUnaryServerInterceptorBuilder interface {
	GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor
	Use(intercepter grpc.UnaryServerInterceptor)
}
