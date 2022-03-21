// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/auth (interfaces: IModularAuthMiddleware)

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockIModularAuthMiddleware is a mock of IModularAuthMiddleware interface.
type MockIModularAuthMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockIModularAuthMiddlewareMockRecorder
}

// MockIModularAuthMiddlewareMockRecorder is the mock recorder for MockIModularAuthMiddleware.
type MockIModularAuthMiddlewareMockRecorder struct {
	mock *MockIModularAuthMiddleware
}

// NewMockIModularAuthMiddleware creates a new mock instance.
func NewMockIModularAuthMiddleware(ctrl *gomock.Controller) *MockIModularAuthMiddleware {
	mock := &MockIModularAuthMiddleware{ctrl: ctrl}
	mock.recorder = &MockIModularAuthMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIModularAuthMiddleware) EXPECT() *MockIModularAuthMiddlewareMockRecorder {
	return m.recorder
}

// GetUnaryServerInterceptor mocks base method.
func (m *MockIModularAuthMiddleware) GetUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnaryServerInterceptor")
	ret0, _ := ret[0].(grpc.UnaryServerInterceptor)
	return ret0
}

// GetUnaryServerInterceptor indicates an expected call of GetUnaryServerInterceptor.
func (mr *MockIModularAuthMiddlewareMockRecorder) GetUnaryServerInterceptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnaryServerInterceptor", reflect.TypeOf((*MockIModularAuthMiddleware)(nil).GetUnaryServerInterceptor))
}