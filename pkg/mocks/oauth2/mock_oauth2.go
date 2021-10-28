// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2 (interfaces: IOauth2)

// Package oauth2 is a generated GoMock package.
package oauth2

import (
	gomock "github.com/golang/mock/gomock"
)

// MockIOauth2 is a mock of IOauth2 interface.
type MockIOauth2 struct {
	ctrl     *gomock.Controller
	recorder *MockIOauth2MockRecorder
}

// MockIOauth2MockRecorder is the mock recorder for MockIOauth2.
type MockIOauth2MockRecorder struct {
	mock *MockIOauth2
}

// NewMockIOauth2 creates a new mock instance.
func NewMockIOauth2(ctrl *gomock.Controller) *MockIOauth2 {
	mock := &MockIOauth2{ctrl: ctrl}
	mock.recorder = &MockIOauth2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOauth2) EXPECT() *MockIOauth2MockRecorder {
	return m.recorder
}