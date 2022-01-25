// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/uuid (interfaces: IKSUID)

// Package uuid is a generated GoMock package.
package uuid

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIKSUID is a mock of IKSUID interface.
type MockIKSUID struct {
	ctrl     *gomock.Controller
	recorder *MockIKSUIDMockRecorder
}

// MockIKSUIDMockRecorder is the mock recorder for MockIKSUID.
type MockIKSUIDMockRecorder struct {
	mock *MockIKSUID
}

// NewMockIKSUID creates a new mock instance.
func NewMockIKSUID(ctrl *gomock.Controller) *MockIKSUID {
	mock := &MockIKSUID{ctrl: ctrl}
	mock.recorder = &MockIKSUIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIKSUID) EXPECT() *MockIKSUIDMockRecorder {
	return m.recorder
}

// UUID mocks base method.
func (m *MockIKSUID) UUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UUID indicates an expected call of UUID.
func (mr *MockIKSUIDMockRecorder) UUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUID", reflect.TypeOf((*MockIKSUID)(nil).UUID))
}
