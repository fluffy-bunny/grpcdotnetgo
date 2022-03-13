// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/singleton (interfaces: ISingleton)

// Package singleton is a generated GoMock package.
package singleton

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISingleton is a mock of ISingleton interface.
type MockISingleton struct {
	ctrl     *gomock.Controller
	recorder *MockISingletonMockRecorder
}

// MockISingletonMockRecorder is the mock recorder for MockISingleton.
type MockISingletonMockRecorder struct {
	mock *MockISingleton
}

// NewMockISingleton creates a new mock instance.
func NewMockISingleton(ctrl *gomock.Controller) *MockISingleton {
	mock := &MockISingleton{ctrl: ctrl}
	mock.recorder = &MockISingletonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISingleton) EXPECT() *MockISingletonMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockISingleton) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockISingletonMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockISingleton)(nil).GetName))
}

// SetName mocks base method.
func (m *MockISingleton) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MockISingletonMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockISingleton)(nil).SetName), arg0)
}
