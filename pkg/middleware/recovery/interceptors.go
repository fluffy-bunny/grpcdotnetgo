// Copyright 2017 David Ackroyd. All Rights Reserved.
// See LICENSE for licensing terms.
// modification of the original to allow access to the full method being called
// this way a user can have more information to make a better decision on how to handle the error.

package grpc_recovery

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecoveryHandlerUnaryFunc is a function that recovers from the panic `p` by returning an abritrary object which could be an `error`.
type RecoveryHandlerUnaryFunc func(fullMethodName string, p interface{}) (resp interface{}, err error)

// RecoveryHandlerStreamFunc is a function that recovers from the panic `p` by returning an `error`.
type RecoveryHandlerStreamFunc func(fullMethodName string, p interface{}) (err error)

// RecoveryHandlerFuncContext is a function that recovers from the panic `p` by returning an abritrary object which could be an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerUnaryFuncContext func(ctx context.Context, fullMethodName string, p interface{}) (resp interface{}, err error)

// RecoveryHandlerStreamFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerStreamFuncContext func(ctx context.Context, fullMethodName string, p interface{}) (err error)

// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				resp, err = recoverFromUnary(ctx, info, r, o.recoveryHandlerUnaryFunc)
				fmt.Println("")
			}
		}()

		resp, err = handler(ctx, req)
		panicked = false
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for panic recovery.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = recoverFromStream(stream.Context(), info, r, o.recoveryHandlerStreamFunc)
			}
		}()

		err = handler(srv, stream)
		panicked = false
		return err
	}
}

func recoverFromUnary(ctx context.Context, info *grpc.UnaryServerInfo, p interface{}, r RecoveryHandlerUnaryFuncContext) (resp interface{}, err error) {
	if r == nil {
		return nil, status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, info.FullMethod, p)
}
func recoverFromStream(ctx context.Context, info *grpc.StreamServerInfo, p interface{}, r RecoveryHandlerStreamFuncContext) (err error) {
	if r == nil {
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, info.FullMethod, p)
}
