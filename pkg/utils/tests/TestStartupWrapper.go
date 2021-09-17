package tests

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/core/types"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// TestStartupWrapper struct
type TestStartupWrapper struct {
	ChildStartup              types.IStartup
	ConfigureServicesOverride func(builder *di.Builder)
}

// NewTestStartupWrapper creates a new TestStartupWrapper
func NewTestStartupWrapper(childStartup types.IStartup, configureServicesOverride func(builder *di.Builder)) *TestStartupWrapper {
	return &TestStartupWrapper{
		ChildStartup:              childStartup,
		ConfigureServicesOverride: configureServicesOverride,
	}
}

// GetConfigOptions wrapper
func (s *TestStartupWrapper) GetConfigOptions() *types.ConfigOptions {
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
func (s *TestStartupWrapper) Configure(unaryServerInterceptorBuilder types.IUnaryServerInterceptorBuilder) {
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
