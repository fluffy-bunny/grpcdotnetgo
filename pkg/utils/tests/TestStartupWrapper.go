package tests

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

type TestStartupWrapperConfig struct {
	InnerStartup          coreContracts.IStartup
	ConfigureServicesHook func(builder *di.Builder)
}

// TestStartupWrapper struct
type TestStartupWrapper struct {
	Config        *TestStartupWrapperConfig
	RootContainer di.Container
}

// NewTestStartupWrapper creates a new TestStartupWrapper
func NewTestStartupWrapper(config *TestStartupWrapperConfig) *TestStartupWrapper {
	return &TestStartupWrapper{
		Config: config,
	}
}

// GetConfigOptions wrapper
func (s *TestStartupWrapper) GetConfigOptions() *coreContracts.ConfigOptions {
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
func (s *TestStartupWrapper) Configure(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder) {
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
func (s *TestStartupWrapper) GetStartupManifest() coreContracts.StartupManifest {
	return coreContracts.StartupManifest{
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
