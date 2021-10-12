package core

import (
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// ICoreConfig ...
type ICoreConfig interface {
	GetPort() int
}

// StartupManifest informational
type StartupManifest struct {
	Name    string
	Version string
}

// IStartup contract
type IStartup interface {
	GetStartupManifest() StartupManifest
	GetConfigOptions() *ConfigOptions
	ConfigureServices(builder *di.Builder)
	Configure(unaryServerInterceptorBuilder IUnaryServerInterceptorBuilder)
	GetPort() int

	RegisterGRPCEndpoints(server *grpc.Server) []interface{}
	SetRootContainer(container di.Container)
	OnPreServerStartup() error
	OnPostServerShutdown()
}

// IUnaryServerInterceptorBuilder ...
type IUnaryServerInterceptorBuilder interface {
	GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor
	Use(intercepter grpc.UnaryServerInterceptor)
}
