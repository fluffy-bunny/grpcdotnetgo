package tests

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// TestStartupWrapper struct
type TestStartupWrapper struct {
	InnerStartup              coreContracts.IStartup
	configureServicesOverride func(builder *di.Builder)
	configureOverride         func(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder)
	RootContainer             di.Container
}

// NewTestStartupWrapper creates a new TestStartupWrapper
func NewTestStartupWrapper(childStartup coreContracts.IStartup,
	configureServicesOverride func(builder *di.Builder),
	configureOverride func(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder)) *TestStartupWrapper {
	return &TestStartupWrapper{
		InnerStartup:              childStartup,
		configureServicesOverride: configureServicesOverride,
		configureOverride:         configureOverride,
	}
}

// GetConfigOptions wrapper
func (s *TestStartupWrapper) GetConfigOptions() *coreContracts.ConfigOptions {
	return s.InnerStartup.GetConfigOptions()
}

// ConfigureServices wrapper
func (s *TestStartupWrapper) ConfigureServices(builder *di.Builder) {
	s.InnerStartup.ConfigureServices(builder)
	if s.configureServicesOverride != nil {
		s.configureServicesOverride(builder)
	}
}

// Configure wrapper
func (s *TestStartupWrapper) Configure(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder) {
	if s.configureOverride != nil {
		s.configureOverride(unaryServerInterceptorBuilder)
	} else {
		s.InnerStartup.Configure(unaryServerInterceptorBuilder)
	}
}

// GetPort wrapper
func (s *TestStartupWrapper) GetPort() int {
	return s.InnerStartup.GetPort()
}

// RegisterGRPCEndpoints wrapper
func (s *TestStartupWrapper) RegisterGRPCEndpoints(server *grpc.Server) []interface{} {
	return s.InnerStartup.RegisterGRPCEndpoints(server)
}

// SetRootContainer wrapper
func (s *TestStartupWrapper) SetRootContainer(container di.Container) {
	s.InnerStartup.SetRootContainer(container)
	s.RootContainer = container
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
	return s.InnerStartup.OnPreServerStartup()
}

// OnPostServerShutdown Wrapper
func (s *TestStartupWrapper) OnPostServerShutdown() {
	s.InnerStartup.OnPostServerShutdown()
}
