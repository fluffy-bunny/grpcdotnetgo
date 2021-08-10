package core

import (
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
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
		serviceProvider servicesServiceProvider.IServiceProvider,
		unaryServerInterceptorBuilder *UnaryServerInterceptorBuilder)
	GetPort() int
	SetPort(port int)
	RegisterGRPCEndpoints(server *grpc.Server)
}

type IUnaryServerInterceptorBuilder interface {
	Use(intercepter grpc.UnaryServerInterceptor)
}
