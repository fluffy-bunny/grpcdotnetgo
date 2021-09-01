// Modified to pass the function name to AuthFunc
//
// From https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/auth/auth.go
// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package logger

import (
	"context"

	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	services_logger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// LoggingUnaryServerInterceptor returns a new unary server interceptors that performs request logging in JSON format
func EnsureContextLoggingUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := log.With().Logger()
		newCtx := logger.WithContext(ctx)
		return handler(newCtx, req)
	}
}

// LoggingUnaryServerInterceptor returns a new unary server interceptors that performs request logging in JSON format
func LoggingUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestContainer := middleware_dicontext.GetRequestContainer(ctx)
		logger := services_logger.GetScopedLoggerFromContainer(requestContainer)
		logger.Debug().
			Interface("request", req).
			Msg("Handling request")
		return handler(ctx, req)
	}
}
