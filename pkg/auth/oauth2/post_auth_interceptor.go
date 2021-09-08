package oauth2

import (
	"context"

	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	services_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"

	"github.com/gogo/status"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func validateAND(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal *ClaimsPrincipal) bool {
	if claimsConfig.AND == nil || len(claimsConfig.AND) > 0 {
		return true
	}
	for _, v := range claimsConfig.AND {
		p, ok := claimsPrincipal.FastMap[v.Type]
		if !ok {
			return false
		}
		_, ok = p[v.Value]
		if !ok {
			return false
		}
	}
	return true
}
func validateOR(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal *ClaimsPrincipal) bool {
	if claimsConfig.OR == nil || len(claimsConfig.OR) > 0 {
		return true
	}

	var found bool = false
	for _, v := range claimsConfig.OR {
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
	return found
}

func FinalAuthVerificationMiddleware(container di.Container) grpc.UnaryServerInterceptor {
	configAccessor := middleware_oidc.GetOIDCConfigAccessorFromContainer(container)
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

		switch claimsPrincipal := ctx.Value(CtxClaimsPrincipalKey).(type) {
		default:
			return permissionDeniedFunc()
		case *ClaimsPrincipal:
			elem, ok := entryPointConfig[info.FullMethod]
			if !ok {
				break
			}
			if !validateAND(elem.ClaimsConfig, claimsPrincipal) {
				return permissionDeniedFunc()
			}
			if !validateOR(elem.ClaimsConfig, claimsPrincipal) {
				return permissionDeniedFunc()
			}

		}
		return handler(ctx, req)
	}
}
