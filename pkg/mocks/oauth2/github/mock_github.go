// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2/github (interfaces: IGithubOAuth2Authenticator)

// Package github is a generated GoMock package.
package github

import (
	context "context"
	reflect "reflect"

	github "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2/github"
	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
)

// MockIGithubOAuth2Authenticator is a mock of IGithubOAuth2Authenticator interface.
type MockIGithubOAuth2Authenticator struct {
	ctrl     *gomock.Controller
	recorder *MockIGithubOAuth2AuthenticatorMockRecorder
}

// MockIGithubOAuth2AuthenticatorMockRecorder is the mock recorder for MockIGithubOAuth2Authenticator.
type MockIGithubOAuth2AuthenticatorMockRecorder struct {
	mock *MockIGithubOAuth2Authenticator
}

// NewMockIGithubOAuth2Authenticator creates a new mock instance.
func NewMockIGithubOAuth2Authenticator(ctrl *gomock.Controller) *MockIGithubOAuth2Authenticator {
	mock := &MockIGithubOAuth2Authenticator{ctrl: ctrl}
	mock.recorder = &MockIGithubOAuth2AuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGithubOAuth2Authenticator) EXPECT() *MockIGithubOAuth2AuthenticatorMockRecorder {
	return m.recorder
}

// AuthCodeURL mocks base method.
func (m *MockIGithubOAuth2Authenticator) AuthCodeURL(arg0 string, arg1 ...oauth2.AuthCodeOption) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthCodeURL", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthCodeURL indicates an expected call of AuthCodeURL.
func (mr *MockIGithubOAuth2AuthenticatorMockRecorder) AuthCodeURL(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthCodeURL", reflect.TypeOf((*MockIGithubOAuth2Authenticator)(nil).AuthCodeURL), varargs...)
}

// Exchange mocks base method.
func (m *MockIGithubOAuth2Authenticator) Exchange(arg0 context.Context, arg1 string, arg2 ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exchange", varargs...)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exchange indicates an expected call of Exchange.
func (mr *MockIGithubOAuth2AuthenticatorMockRecorder) Exchange(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*MockIGithubOAuth2Authenticator)(nil).Exchange), varargs...)
}

// GetTokenSource mocks base method.
func (m *MockIGithubOAuth2Authenticator) GetTokenSource(arg0 context.Context, arg1 *oauth2.Token) oauth2.TokenSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenSource", arg0, arg1)
	ret0, _ := ret[0].(oauth2.TokenSource)
	return ret0
}

// GetTokenSource indicates an expected call of GetTokenSource.
func (mr *MockIGithubOAuth2AuthenticatorMockRecorder) GetTokenSource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenSource", reflect.TypeOf((*MockIGithubOAuth2Authenticator)(nil).GetTokenSource), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockIGithubOAuth2Authenticator) GetUser(arg0 *oauth2.Token) (*github.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*github.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIGithubOAuth2AuthenticatorMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIGithubOAuth2Authenticator)(nil).GetUser), arg0)
}
