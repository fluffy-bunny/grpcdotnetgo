package oidc

import (
	"context"

	middleware_grpc_auth "github.com/fluffy-bunny/grpcdotnetgo/middleware/auth"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/middleware/oidc"
	services_logger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	services_oidc "github.com/fluffy-bunny/grpcdotnetgo/services/oidc"
	"github.com/gogo/status"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
)

type service struct {
	Logger     services_logger.ILogger
	Storage    services_oidc.IOidcBackgroundStorage
	OIDCConfig middleware_oidc.IOIDCConfig
}

func (s *service) GetAuthFuncStream() middleware_grpc_auth.AuthFuncStream {
	return nil
}

func (s *service) GetAuthFuncUnary() middleware_grpc_auth.AuthFuncUnary {
	return func(ctx context.Context, fullMethodName string) (context.Context, interface{}, error) {
		disco := s.Storage.AtomicGet()
		if disco == nil {
			s.Logger.Error().Msg("Discovery Document is nil, make sure the background job to fetch it is in place.")
			return ctx, nil, status.Error(codes.Unauthenticated, "Unauthorized")
		}
		entryPoints := s.OIDCConfig.GetEntryPoints()
		log.Info().Interface("config", entryPoints).Send()
		_, ok := entryPoints[fullMethodName]
		if ok {
			token, err := middleware_grpc_auth.AuthFromMD(ctx, "bearer")
			if err != nil || token == "" {

				return ctx, nil, status.Error(codes.Unauthenticated, "Unauthorized")
			}
		}

		return ctx, nil, nil
	}
}
