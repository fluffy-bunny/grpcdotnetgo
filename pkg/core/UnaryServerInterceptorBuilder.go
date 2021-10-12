package core

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	"google.golang.org/grpc"
)

// NewUnaryServerInterceptorBuilder helper
func NewUnaryServerInterceptorBuilder() coreContracts.IUnaryServerInterceptorBuilder {
	return &UnaryServerInterceptorBuilder{}
}

// UnaryServerInterceptorBuilder struct
type UnaryServerInterceptorBuilder struct {
	UnaryServerInterceptors []grpc.UnaryServerInterceptor
}

// Use helper
func (s *UnaryServerInterceptorBuilder) Use(interceptor grpc.UnaryServerInterceptor) {
	s.UnaryServerInterceptors = append(s.UnaryServerInterceptors, interceptor)
}

// GetUnaryServerInterceptors helper
func (s *UnaryServerInterceptorBuilder) GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	return s.UnaryServerInterceptors
}
