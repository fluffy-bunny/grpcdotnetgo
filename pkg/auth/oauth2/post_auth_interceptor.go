package oauth2

import (
	"context"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gogo/status"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func validate(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if !validateAND(claimsConfig, claimsPrincipal) {
		return false
	}
	if !validateOR(claimsConfig, claimsPrincipal) {
		return false
	}
	return true
}
func validateAND(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if claimsConfig.AND == nil || len(claimsConfig.AND) == 0 {
		return true
	}
	for _, v := range claimsConfig.AND {
		if !claimsPrincipal.HasClaim(v) {
			return false
		}
	}
	return true
}
func validateOR(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if claimsConfig.OR == nil || len(claimsConfig.OR) == 0 {
		return true
	}
	for _, v := range claimsConfig.OR {
		if claimsPrincipal.HasClaim(v) {
			return true
		}
	}
	return false
}

// FinalAuthVerificationMiddleware evaluates the claims principal
func FinalAuthVerificationMiddleware(container di.Container) grpc.UnaryServerInterceptor {
	configAccessor := middleware_oidc.GetOIDCConfigAccessorFromContainer(container)
	entryPointConfig := configAccessor.GetOIDCConfig().GetEntryPoints()
	log.Info().Interface("entryPointConfig", entryPointConfig).Send()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestContainer := middleware_dicontext.GetRequestContainer(ctx)
		if requestContainer != nil {
			claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)
			logger := loggerContracts.GetILoggerFromContainer(requestContainer)
			loggerZ := logger.GetLogger()
			subLogger := loggerZ.With().Str("FullMethod", info.FullMethod).Logger()

			permissionDeniedFunc := func() (interface{}, error) {
				logger.DebugL(&subLogger).Msg("")
				return nil, status.Errorf(codes.PermissionDenied, "permission denied")
			}
			elem, ok := entryPointConfig[info.FullMethod]
			if !ok {
				// we don't have an entry so it is valid
				// TODO: Add in security option that must have an entry even if the AND and OR are empty.  That way
				// We have proof someone purposefully wanted it with no validation
				// return permissionDeniedFunc()
			} else {
				if !validate(elem.ClaimsConfig, claimsPrincipal) {
					return permissionDeniedFunc()
				}
			}
		}
		return handler(ctx, req)
	}
}
