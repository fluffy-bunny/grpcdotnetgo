package tests

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// TestStartupWrapper struct
type TestStartupWrapper struct {
	ChildStartup              coreContracts.IStartup
	ConfigureServicesOverride func(builder *di.Builder)
}

// NewTestStartupWrapper creates a new TestStartupWrapper
func NewTestStartupWrapper(childStartup coreContracts.IStartup, configureServicesOverride func(builder *di.Builder)) *TestStartupWrapper {
	return &TestStartupWrapper{
		ChildStartup:              childStartup,
		ConfigureServicesOverride: configureServicesOverride,
	}
}

// GetConfigOptions wrapper
func (s *TestStartupWrapper) GetConfigOptions() *coreContracts.ConfigOptions {
	return s.ChildStartup.GetConfigOptions()
}

// ConfigureServices wrapper
func (s *TestStartupWrapper) ConfigureServices(builder *di.Builder) {
	s.ChildStartup.ConfigureServices(builder)
	if s.ConfigureServicesOverride != nil {
		s.ConfigureServicesOverride(builder)
	}
}

// Configure wrapper
func (s *TestStartupWrapper) Configure(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder) {
	s.Configure(unaryServerInterceptorBuilder)
}

// GetPort wrapper
func (s *TestStartupWrapper) GetPort() int {
	return s.ChildStartup.GetPort()
}

// RegisterGRPCEndpoints wrapper
func (s *TestStartupWrapper) RegisterGRPCEndpoints(server *grpc.Server) []interface{} {
	return s.ChildStartup.RegisterGRPCEndpoints(server)
}

// SetRootContainer wrapper
func (s *TestStartupWrapper) SetRootContainer(container di.Container) {
	s.SetRootContainer(container)
}

// GetStartupManifest wrapper
func (s *TestStartupWrapper) GetStartupManifest() coreContracts.StartupManifest {
	return coreContracts.StartupManifest{
		Name:    "test",
		Version: "test.1",
	}
}

// OnPreServerStartup wrapper
func (s *TestStartupWrapper) OnPreServerStartup() error {
	return nil
}

// OnPostServerShutdown Wrapper
func (s *TestStartupWrapper) OnPostServerShutdown() {}
