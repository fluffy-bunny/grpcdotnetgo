package tests

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// TestStartupWrapper struct
type TestStartupWrapper struct {
	innerStartup              coreContracts.IStartup
	configureServicesOverride func(builder *di.Builder)
	configureOverride         func(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder)
}

// NewTestStartupWrapper creates a new TestStartupWrapper
func NewTestStartupWrapper(childStartup coreContracts.IStartup,
	configureServicesOverride func(builder *di.Builder),
	configureOverride func(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder)) coreContracts.IStartup {
	return &TestStartupWrapper{
		innerStartup:              childStartup,
		configureServicesOverride: configureServicesOverride,
		configureOverride:         configureOverride,
	}
}

// GetConfigOptions wrapper
func (s *TestStartupWrapper) GetConfigOptions() *coreContracts.ConfigOptions {
	return s.innerStartup.GetConfigOptions()
}

// ConfigureServices wrapper
func (s *TestStartupWrapper) ConfigureServices(builder *di.Builder) {
	s.innerStartup.ConfigureServices(builder)
	if s.configureServicesOverride != nil {
		s.configureServicesOverride(builder)
	}
}

// Configure wrapper
func (s *TestStartupWrapper) Configure(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder) {
	if s.configureOverride != nil {
		s.configureOverride(unaryServerInterceptorBuilder)
	} else {
		s.innerStartup.Configure(unaryServerInterceptorBuilder)
	}
}

// GetPort wrapper
func (s *TestStartupWrapper) GetPort() int {
	return s.innerStartup.GetPort()
}

// RegisterGRPCEndpoints wrapper
func (s *TestStartupWrapper) RegisterGRPCEndpoints(server *grpc.Server) []interface{} {
	return s.innerStartup.RegisterGRPCEndpoints(server)
}

// SetRootContainer wrapper
func (s *TestStartupWrapper) SetRootContainer(container di.Container) {
	s.innerStartup.SetRootContainer(container)
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
	return s.innerStartup.OnPreServerStartup()
}

// OnPostServerShutdown Wrapper
func (s *TestStartupWrapper) OnPostServerShutdown() {
	s.innerStartup.OnPostServerShutdown()
}
