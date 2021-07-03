package core

import "google.golang.org/grpc"

type UnaryServerInterceptorBuilder struct {
	UnaryServerInterceptors []grpc.UnaryServerInterceptor
}

func (s *UnaryServerInterceptorBuilder) Use(interceptor grpc.UnaryServerInterceptor) {
	s.UnaryServerInterceptors = append(s.UnaryServerInterceptors, interceptor)
}
