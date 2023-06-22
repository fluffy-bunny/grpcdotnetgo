package tests

import (
	"context"

	contracts_core "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

type TestStartupWrapperConfig struct {
	InnerStartup          contracts_core.IStartup
	ConfigureServicesHook func(builder *di.Builder)
}

// TestStartupWrapper struct
type TestStartupWrapper struct {
	contracts_core.UnimplementedStartup

	Config        *TestStartupWrapperConfig
	RootContainer di.Container
}

// NewTestStartupWrapper creates a new TestStartupWrapper
func NewTestStartupWrapper(config *TestStartupWrapperConfig) *TestStartupWrapper {
	return &TestStartupWrapper{
		Config: config,
	}
}
func (s *TestStartupWrapper) SetContext(ctx context.Context) {
	s.UnimplementedStartup.SetContext(ctx)
	s.Config.InnerStartup.SetContext(ctx)
}

// GetConfigOptions wrapper
func (s *TestStartupWrapper) GetConfigOptions() *contracts_core.ConfigOptions {
	return s.Config.InnerStartup.GetConfigOptions()
}

// ConfigureServices wrapper
func (s *TestStartupWrapper) ConfigureServices(builder *di.Builder) {
	s.Config.InnerStartup.ConfigureServices(builder)
	if s.Config.ConfigureServicesHook != nil {
		s.Config.ConfigureServicesHook(builder)
	}
}

// Configure wrapper
func (s *TestStartupWrapper) Configure(unaryServerInterceptorBuilder contracts_core.IUnaryServerInterceptorBuilder) {
	s.Config.InnerStartup.Configure(unaryServerInterceptorBuilder)
}

// GetPort wrapper
func (s *TestStartupWrapper) GetPort() int {
	return s.Config.InnerStartup.GetPort()
}

// RegisterGRPCEndpoints wrapper
func (s *TestStartupWrapper) RegisterGRPCEndpoints(server *grpc.Server) []interface{} {
	return s.Config.InnerStartup.RegisterGRPCEndpoints(server)
}

// SetRootContainer wrapper
func (s *TestStartupWrapper) SetRootContainer(container di.Container) {
	s.Config.InnerStartup.SetRootContainer(container)
	s.RootContainer = container
}

// GetStartupManifest wrapper
func (s *TestStartupWrapper) GetStartupManifest() contracts_core.StartupManifest {
	return contracts_core.StartupManifest{
		Name:    "test",
		Version: "test.1",
	}
}

// OnPreServerStartup wrapper
func (s *TestStartupWrapper) OnPreServerStartup() error {
	return s.Config.InnerStartup.OnPreServerStartup()
}

// OnPostServerShutdown Wrapper
func (s *TestStartupWrapper) OnPostServerShutdown() {
	s.Config.InnerStartup.OnPostServerShutdown()
}
