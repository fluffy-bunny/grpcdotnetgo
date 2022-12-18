// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core (interfaces: ICoreConfig,IStartup,IUnaryServerInterceptorBuilder)

// Package core is a generated GoMock package.
package core

import (
	reflect "reflect"

	core "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockICoreConfig is a mock of ICoreConfig interface.
type MockICoreConfig struct {
	ctrl     *gomock.Controller
	recorder *MockICoreConfigMockRecorder
}

// MockICoreConfigMockRecorder is the mock recorder for MockICoreConfig.
type MockICoreConfigMockRecorder struct {
	mock *MockICoreConfig
}

// NewMockICoreConfig creates a new mock instance.
func NewMockICoreConfig(ctrl *gomock.Controller) *MockICoreConfig {
	mock := &MockICoreConfig{ctrl: ctrl}
	mock.recorder = &MockICoreConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICoreConfig) EXPECT() *MockICoreConfigMockRecorder {
	return m.recorder
}

// GetPort mocks base method.
func (m *MockICoreConfig) GetPort() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPort indicates an expected call of GetPort.
func (mr *MockICoreConfigMockRecorder) GetPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockICoreConfig)(nil).GetPort))
}

// MockIStartup is a mock of IStartup interface.
type MockIStartup struct {
	ctrl     *gomock.Controller
	recorder *MockIStartupMockRecorder
}

// MockIStartupMockRecorder is the mock recorder for MockIStartup.
type MockIStartupMockRecorder struct {
	mock *MockIStartup
}

// NewMockIStartup creates a new mock instance.
func NewMockIStartup(ctrl *gomock.Controller) *MockIStartup {
	mock := &MockIStartup{ctrl: ctrl}
	mock.recorder = &MockIStartupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStartup) EXPECT() *MockIStartupMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *MockIStartup) Configure(arg0 core.IUnaryServerInterceptorBuilder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Configure", arg0)
}

// Configure indicates an expected call of Configure.
func (mr *MockIStartupMockRecorder) Configure(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockIStartup)(nil).Configure), arg0)
}

// ConfigureServices mocks base method.
func (m *MockIStartup) ConfigureServices(arg0 *di.Builder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConfigureServices", arg0)
}

// ConfigureServices indicates an expected call of ConfigureServices.
func (mr *MockIStartupMockRecorder) ConfigureServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureServices", reflect.TypeOf((*MockIStartup)(nil).ConfigureServices), arg0)
}

// GetConfigOptions mocks base method.
func (m *MockIStartup) GetConfigOptions() *core.ConfigOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigOptions")
	ret0, _ := ret[0].(*core.ConfigOptions)
	return ret0
}

// GetConfigOptions indicates an expected call of GetConfigOptions.
func (mr *MockIStartupMockRecorder) GetConfigOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigOptions", reflect.TypeOf((*MockIStartup)(nil).GetConfigOptions))
}

// GetPort mocks base method.
func (m *MockIStartup) GetPort() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPort indicates an expected call of GetPort.
func (mr *MockIStartupMockRecorder) GetPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockIStartup)(nil).GetPort))
}

// GetStartupManifest mocks base method.
func (m *MockIStartup) GetStartupManifest() core.StartupManifest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartupManifest")
	ret0, _ := ret[0].(core.StartupManifest)
	return ret0
}

// GetStartupManifest indicates an expected call of GetStartupManifest.
func (mr *MockIStartupMockRecorder) GetStartupManifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartupManifest", reflect.TypeOf((*MockIStartup)(nil).GetStartupManifest))
}

// OnPostServerShutdown mocks base method.
func (m *MockIStartup) OnPostServerShutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPostServerShutdown")
}

// OnPostServerShutdown indicates an expected call of OnPostServerShutdown.
func (mr *MockIStartupMockRecorder) OnPostServerShutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPostServerShutdown", reflect.TypeOf((*MockIStartup)(nil).OnPostServerShutdown))
}

// OnPreServerStartup mocks base method.
func (m *MockIStartup) OnPreServerStartup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnPreServerStartup")
	ret0, _ := ret[0].(error)
	return ret0
}

// OnPreServerStartup indicates an expected call of OnPreServerStartup.
func (mr *MockIStartupMockRecorder) OnPreServerStartup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPreServerStartup", reflect.TypeOf((*MockIStartup)(nil).OnPreServerStartup))
}

// RegisterGRPCEndpoints mocks base method.
func (m *MockIStartup) RegisterGRPCEndpoints(arg0 *grpc.Server) []interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterGRPCEndpoints", arg0)
	ret0, _ := ret[0].([]interface{})
	return ret0
}

// RegisterGRPCEndpoints indicates an expected call of RegisterGRPCEndpoints.
func (mr *MockIStartupMockRecorder) RegisterGRPCEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCEndpoints", reflect.TypeOf((*MockIStartup)(nil).RegisterGRPCEndpoints), arg0)
}

// SetRootContainer mocks base method.
func (m *MockIStartup) SetRootContainer(arg0 di.Container) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRootContainer", arg0)
}

// SetRootContainer indicates an expected call of SetRootContainer.
func (mr *MockIStartupMockRecorder) SetRootContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRootContainer", reflect.TypeOf((*MockIStartup)(nil).SetRootContainer), arg0)
}

// mustEmbedUnimplementedStartup mocks base method.
func (m *MockIStartup) mustEmbedUnimplementedStartup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStartup")
}

// mustEmbedUnimplementedStartup indicates an expected call of mustEmbedUnimplementedStartup.
func (mr *MockIStartupMockRecorder) mustEmbedUnimplementedStartup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStartup", reflect.TypeOf((*MockIStartup)(nil).mustEmbedUnimplementedStartup))
}

// MockIUnaryServerInterceptorBuilder is a mock of IUnaryServerInterceptorBuilder interface.
type MockIUnaryServerInterceptorBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockIUnaryServerInterceptorBuilderMockRecorder
}

// MockIUnaryServerInterceptorBuilderMockRecorder is the mock recorder for MockIUnaryServerInterceptorBuilder.
type MockIUnaryServerInterceptorBuilderMockRecorder struct {
	mock *MockIUnaryServerInterceptorBuilder
}

// NewMockIUnaryServerInterceptorBuilder creates a new mock instance.
func NewMockIUnaryServerInterceptorBuilder(ctrl *gomock.Controller) *MockIUnaryServerInterceptorBuilder {
	mock := &MockIUnaryServerInterceptorBuilder{ctrl: ctrl}
	mock.recorder = &MockIUnaryServerInterceptorBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUnaryServerInterceptorBuilder) EXPECT() *MockIUnaryServerInterceptorBuilderMockRecorder {
	return m.recorder
}

// GetUnaryServerInterceptors mocks base method.
func (m *MockIUnaryServerInterceptorBuilder) GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnaryServerInterceptors")
	ret0, _ := ret[0].([]grpc.UnaryServerInterceptor)
	return ret0
}

// GetUnaryServerInterceptors indicates an expected call of GetUnaryServerInterceptors.
func (mr *MockIUnaryServerInterceptorBuilderMockRecorder) GetUnaryServerInterceptors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnaryServerInterceptors", reflect.TypeOf((*MockIUnaryServerInterceptorBuilder)(nil).GetUnaryServerInterceptors))
}

// Use mocks base method.
func (m *MockIUnaryServerInterceptorBuilder) Use(arg0 grpc.UnaryServerInterceptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Use", arg0)
}

// Use indicates an expected call of Use.
func (mr *MockIUnaryServerInterceptorBuilderMockRecorder) Use(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockIUnaryServerInterceptorBuilder)(nil).Use), arg0)
}
