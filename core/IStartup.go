package core

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

type ICoreConfig interface {
	GetPort() int
}
type ConfigOptions struct {
	Destination    interface{}
	RootConfigYaml []byte
	ConfigPath     string
}
type IStartup interface {
	GetConfigOptions() *ConfigOptions
	ConfigureServices(builder *di.Builder)
	Configure(
		container di.Container,
		unaryServerInterceptorBuilder *UnaryServerInterceptorBuilder)
	GetPort() int
	SetPort(port int)
	RegisterGRPCEndpoints(server *grpc.Server)
}

type IUnaryServerInterceptorBuilder interface {
	Use(intercepter grpc.UnaryServerInterceptor)
}
