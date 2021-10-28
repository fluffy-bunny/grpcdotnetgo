package core

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ICoreConfig,IStartup,IUnaryServerInterceptorBuilder"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ICoreConfig,IStartup,IUnaryServerInterceptorBuilder

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
