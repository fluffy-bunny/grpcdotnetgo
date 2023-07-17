package core

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ICoreConfig,IStartup,IUnaryServerInterceptorBuilder"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ICoreConfig,IStartup,IUnaryServerInterceptorBuilder

import (
	"context"

	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// ICoreConfig ...
type ICoreConfig interface {
	GetPort() int
}

// StartupManifest informational
type StartupManifest struct {
	Name               string `json:"name" mapstructure:"NAME"`
	Version            string `json:"version" mapstructure:"VERSION"`
	Port               int    `json:"port" mapstructure:"PORT"`
	GRPCGatewayEnabled bool   `json:"grpcGatewayEnabled" mapstructure:"GRPC_GATEWAY_ENABLED"`
	RESTPort           int    `json:"restPort" mapstructure:"REST_PORT"`
}

// UnimplementedStartup helper ...
type UnimplementedStartup struct {
	Context context.Context
}

func (u UnimplementedStartup) mustEmbedUnimplementedStartup() {

}
func (u *UnimplementedStartup) SetContext(ctx context.Context) {
	u.Context = ctx
}

// RegisterGRPCEndpoints legacy
func (u UnimplementedStartup) RegisterGRPCEndpoints(_ *grpc.Server) []interface{} {
	return nil
}

// OnPreServerStartup ...
func (u UnimplementedStartup) OnPreServerStartup() error { return nil }

// OnPostServerShutdown ...
func (u UnimplementedStartup) OnPostServerShutdown() {}

// GetPort ...
func (u UnimplementedStartup) GetPort() int {
	return 0
}

// IStartup contract
type IStartup interface {
	mustEmbedUnimplementedStartup()
	GetStartupManifest() StartupManifest
	GetConfigOptions() *ConfigOptions
	ConfigureServices(builder *di.Builder)
	Configure(unaryServerInterceptorBuilder IUnaryServerInterceptorBuilder)
	// GetPort returns the port number.
	// Deprecated: pass the port in the StartupManifest
	GetPort() int

	// RegisterGRPCEndpoints registers the grpc endpoints with the server.
	// Deprecated: Server endpoints are now added via the DI.  i.e. AddGreeterEndpointRegistration(...)
	RegisterGRPCEndpoints(server *grpc.Server) []interface{}
	SetRootContainer(container di.Container)
	OnPreServerStartup() error
	OnPostServerShutdown()
	SetContext(ctx context.Context)
}

// IUnaryServerInterceptorBuilder ...
type IUnaryServerInterceptorBuilder interface {
	GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor
	Use(intercepter grpc.UnaryServerInterceptor)
}
