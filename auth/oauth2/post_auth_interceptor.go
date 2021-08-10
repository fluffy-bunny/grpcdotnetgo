package oauth2

import (
	"context"

	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/middleware/dicontext"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/middleware/oidc"
	services_logger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	"github.com/gogo/status"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func FinalAuthVerificationMiddleware(serviceProvider servicesServiceProvider.IServiceProvider) grpc.UnaryServerInterceptor {
	configAccessor := middleware_oidc.GetOIDCConfigAccessorFromContainer(serviceProvider.GetContainer())
	entryPointConfig := configAccessor.GetOIDCConfig().GetEntryPoints()
	log.Info().Interface("entryPointConfig", entryPointConfig).Send()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestContainer := middleware_dicontext.GetRequestContainer(ctx)
		logger := services_logger.GetScopedLoggerFromContainer(requestContainer)
		loggerZ := logger.GetLogger()
		subLogger := loggerZ.With().Str("FullMethod", info.FullMethod).Logger()

		permissionDeniedFunc := func() (interface{}, error) {
			logger.DebugL(&subLogger).Msg("")
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
		data := ctx.Value(CtxClaimsPrincipalKey)
		if data == nil {
			return permissionDeniedFunc()
		}

		claimsPrincipal := data.(*ClaimsPrincipal)
		elem, ok := entryPointConfig[info.FullMethod]
		if ok {
			for _, v := range elem.ClaimsConfig.AND {
				p, ok := claimsPrincipal.FastMap[v.Type]
				if !ok {
					return permissionDeniedFunc()
				}
				_, ok = p[v.Value]
				if !ok {
					return permissionDeniedFunc()
				}
			}

			if elem.ClaimsConfig.OR != nil && len(elem.ClaimsConfig.OR) > 0 {
				var found bool = false
				for _, v := range elem.ClaimsConfig.OR {
					p, ok := claimsPrincipal.FastMap[v.Type]
					if !ok {
						continue
					}
					_, ok = p[v.Value]
					if !ok {
						continue
					}
					found = true
					break
				}
				if !found {
					return permissionDeniedFunc()
				}
			}

		}

		return handler(ctx, req)
	}

}
