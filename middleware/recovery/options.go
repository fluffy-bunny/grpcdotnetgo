// Copyright 2017 David Ackroyd. All Rights Reserved.
// See LICENSE for licensing terms.
// original

package grpc_recovery

import "context"

var (
	defaultOptions = &options{
		recoveryHandlerUnaryFunc:  nil,
		recoveryHandlerStreamFunc: nil,
	}
)

type options struct {
	recoveryHandlerUnaryFunc  RecoveryHandlerUnaryFuncContext
	recoveryHandlerStreamFunc RecoveryHandlerStreamFuncContext
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type Option func(*options)

// WithRecoveryHandlerUnary customizes the function for recovering from a panic.
func WithRecoveryHandlerUnary(f RecoveryHandlerUnaryFunc) Option {
	return func(o *options) {
		o.recoveryHandlerUnaryFunc = RecoveryHandlerUnaryFuncContext(func(ctx context.Context, fullMethodName string, p interface{}) (interface{}, error) {
			return f(fullMethodName, p)
		})
	}
}

// WithRecoveryHandlerUnaryContext customizes the function for recovering from a panic.
func WithRecoveryHandlerUnaryContext(f RecoveryHandlerUnaryFuncContext) Option {
	return func(o *options) {
		o.recoveryHandlerUnaryFunc = f
	}
}

// WithRecoveryHandlerStream customizes the function for recovering from a panic.
func WithRecoveryHandlerStream(f RecoveryHandlerStreamFunc) Option {
	return func(o *options) {
		o.recoveryHandlerStreamFunc = RecoveryHandlerStreamFuncContext(func(ctx context.Context, fullMethodName string, p interface{}) error {
			return f(fullMethodName, p)
		})
	}
}

// WithRecoveryHandlerUnaryStreamContext customizes the function for recovering from a panic.
func WithRecoveryHandlerUnaryStreamContext(f RecoveryHandlerStreamFuncContext) Option {
	return func(o *options) {
		o.recoveryHandlerStreamFunc = f
	}
}
