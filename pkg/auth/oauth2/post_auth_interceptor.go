package oauth2

import (
	"context"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gogo/status"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func validate(logEvent *zerolog.Event, claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if !validateAND(claimsConfig, claimsPrincipal) {
		logEvent.Msg("AND validation failed")
		return false
	}
	if !validateANDTYPE(claimsConfig, claimsPrincipal) {
		logEvent.Msg("AND validation failed")
		return false
	}
	if !validateOR(claimsConfig, claimsPrincipal) {
		logEvent.Msg("OR validation failed")
		return false
	}
	if !validateORTYPE(claimsConfig, claimsPrincipal) {
		logEvent.Msg("OR validation failed")
		return false
	}
	return true
}
func validateAND(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if utils.IsEmptyOrNil(claimsConfig.AND) {
		return true
	}
	for _, v := range claimsConfig.AND {
		if !claimsPrincipal.HasClaim(v) {
			return false
		}
	}
	return true
}
func validateANDTYPE(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if utils.IsEmptyOrNil(claimsConfig.ANDTYPE) {
		return true
	}
	for _, v := range claimsConfig.ANDTYPE {
		if !claimsPrincipal.HasClaimType(v) {
			return false
		}
	}
	return true
}

func validateOR(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if utils.IsEmptyOrNil(claimsConfig.OR) {
		return true
	}
	for _, v := range claimsConfig.OR {
		if claimsPrincipal.HasClaim(v) {
			return true
		}
	}
	return false
}

func validateORTYPE(claimsConfig middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if utils.IsEmptyOrNil(claimsConfig.ORTYPE) {
		return true
	}
	for _, v := range claimsConfig.ORTYPE {
		if claimsPrincipal.HasClaimType(v) {
			return true
		}
	}
	return false
}

// FinalAuthVerificationMiddleware evaluates the claims principal
func FinalAuthVerificationMiddleware(container di.Container) grpc.UnaryServerInterceptor {
	configAccessor := middleware_oidc.GetOIDCConfigAccessorFromContainer(container)
	entryPointConfig := configAccessor.GetOIDCConfig().GetEntryPoints()
	return FinalAuthVerificationMiddlewareUsingClaimsMap(entryPointConfig)
}

// FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption evaluates the claims principal
func FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption(grpcEntrypointClaimsMap map[string]middleware_oidc.EntryPointConfig, enableZeroTrust bool) grpc.UnaryServerInterceptor {

	log.Info().Interface("entryPointConfig", grpcEntrypointClaimsMap).Send()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestContainer := middleware_dicontext.GetRequestContainer(ctx)
		var subLogger zerolog.Logger
		if requestContainer != nil {
			logger := loggerContracts.GetILoggerFromContainer(requestContainer)
			loggerZ := logger.GetLogger()
			subLogger = loggerZ.With().
				Bool("enableZeroTrust", enableZeroTrust).
				Str("FullMethod", info.FullMethod).
				Logger()
		} else {
			subLogger = log.With().Bool("enableZeroTrust", enableZeroTrust).
				Str("FullMethod", info.FullMethod).
				Logger()
		}

		debugEvent := subLogger.Debug()
		debugEvent.Msg("FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption Enter")
		defer debugEvent.Msg("FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption Exit")
		if requestContainer != nil {
			claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)

			permissionDeniedFunc := func() (interface{}, error) {
				debugEvent.Msg("Permission denied")
				return nil, status.Errorf(codes.PermissionDenied, "permission denied")
			}
			elem, ok := grpcEntrypointClaimsMap[info.FullMethod]
			if !ok && enableZeroTrust {
				debugEvent.Msg("FullMethod not found in entrypoint claims map")
				return permissionDeniedFunc()
			}
			if !validate(debugEvent, elem.ClaimsConfig, claimsPrincipal) {
				debugEvent.Msg("ClaimsConfig validation failed")
				return permissionDeniedFunc()
			}
		}
		return handler(ctx, req)
	}
}

// FinalAuthVerificationMiddlewareUsingClaimsMap evaluates the claims principal
func FinalAuthVerificationMiddlewareUsingClaimsMap(grpcEntrypointClaimsMap map[string]middleware_oidc.EntryPointConfig) grpc.UnaryServerInterceptor {
	return FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption(grpcEntrypointClaimsMap, false)
}

// FinalAuthVerificationMiddlewareUsingClaimsMapWithZeroTrust evaluates the claims principal
func FinalAuthVerificationMiddlewareUsingClaimsMapWithZeroTrust(grpcEntrypointClaimsMap map[string]middleware_oidc.EntryPointConfig) grpc.UnaryServerInterceptor {
	return FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption(grpcEntrypointClaimsMap, true)
}
