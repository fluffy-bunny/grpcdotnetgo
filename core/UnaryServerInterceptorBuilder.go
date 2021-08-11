package core

import (
	"github.com/fluffy-bunny/grpcdotnetgo/core/types"
	"google.golang.org/grpc"
)

func NewUnaryServerInterceptorBuilder() types.IUnaryServerInterceptorBuilder {
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
