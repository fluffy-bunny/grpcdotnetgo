package core

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	"google.golang.org/grpc"
)

func NewUnaryServerInterceptorBuilder() coreContracts.IUnaryServerInterceptorBuilder {
	return &UnaryServerInterceptorBuilder{}
}

type UnaryServerInterceptorBuilder struct {
	UnaryServerInterceptors []grpc.UnaryServerInterceptor
}

func (s *UnaryServerInterceptorBuilder) Use(interceptor grpc.UnaryServerInterceptor) {
	s.UnaryServerInterceptors = append(s.UnaryServerInterceptors, interceptor)
}
func (s *UnaryServerInterceptorBuilder) GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	return s.UnaryServerInterceptors
}
