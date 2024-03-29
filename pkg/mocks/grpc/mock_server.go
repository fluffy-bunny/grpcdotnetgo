// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/grpc (interfaces: IServiceEndpointRegistration)

// Package grpc is a generated GoMock package.
package grpc

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
)

// MockIServiceEndpointRegistration is a mock of IServiceEndpointRegistration interface.
type MockIServiceEndpointRegistration struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceEndpointRegistrationMockRecorder
}

// MockIServiceEndpointRegistrationMockRecorder is the mock recorder for MockIServiceEndpointRegistration.
type MockIServiceEndpointRegistrationMockRecorder struct {
	mock *MockIServiceEndpointRegistration
}

// NewMockIServiceEndpointRegistration creates a new mock instance.
func NewMockIServiceEndpointRegistration(ctrl *gomock.Controller) *MockIServiceEndpointRegistration {
	mock := &MockIServiceEndpointRegistration{ctrl: ctrl}
	mock.recorder = &MockIServiceEndpointRegistrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIServiceEndpointRegistration) EXPECT() *MockIServiceEndpointRegistrationMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockIServiceEndpointRegistration) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockIServiceEndpointRegistrationMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIServiceEndpointRegistration)(nil).GetName))
}

// GetNewClient mocks base method.
func (m *MockIServiceEndpointRegistration) GetNewClient(arg0 grpc.ClientConnInterface) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewClient", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetNewClient indicates an expected call of GetNewClient.
func (mr *MockIServiceEndpointRegistrationMockRecorder) GetNewClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewClient", reflect.TypeOf((*MockIServiceEndpointRegistration)(nil).GetNewClient), arg0)
}

// RegisterEndpoint mocks base method.
func (m *MockIServiceEndpointRegistration) RegisterEndpoint(arg0 *grpc.Server) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterEndpoint", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// RegisterEndpoint indicates an expected call of RegisterEndpoint.
func (mr *MockIServiceEndpointRegistrationMockRecorder) RegisterEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterEndpoint", reflect.TypeOf((*MockIServiceEndpointRegistration)(nil).RegisterEndpoint), arg0)
}

// RegisterGatewayHandler mocks base method.
func (m *MockIServiceEndpointRegistration) RegisterGatewayHandler(arg0 *runtime.ServeMux, arg1 *grpc.ClientConn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGatewayHandler", arg0, arg1)
}

// RegisterGatewayHandler indicates an expected call of RegisterGatewayHandler.
func (mr *MockIServiceEndpointRegistrationMockRecorder) RegisterGatewayHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGatewayHandler", reflect.TypeOf((*MockIServiceEndpointRegistration)(nil).RegisterGatewayHandler), arg0, arg1)
}
