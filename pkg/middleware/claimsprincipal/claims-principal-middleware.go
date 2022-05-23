package claimsprincipal

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

func Validate(logger *zerolog.Logger, claimsConfig *middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if !validateAND(claimsConfig, claimsPrincipal) {
		logger.Debug().Msg("AND validation failed")
		return false
	}

	if !validateOR(claimsConfig, claimsPrincipal) {
		logger.Debug().Msg("OR validation failed")
		return false
	}
	if claimsConfig.Child != nil {
		return Validate(logger, claimsConfig.Child, claimsPrincipal)
	}
	return true
}
func validateAND(claimsConfig *middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if utils.IsEmptyOrNil(claimsConfig.AND) {
		return true
	}

	if !utils.IsEmptyOrNil(claimsConfig.AND) {
		for _, v := range claimsConfig.AND {
			if v.Directive == middleware_oidc.ClaimTypeAndValue {
				if !claimsPrincipal.HasClaim(v.Claim) {
					return false
				}
			}
			if v.Directive == middleware_oidc.ClaimType {
				if !claimsPrincipal.HasClaimType(v.Claim.Type) {
					return false
				}
			}
		}
	}

	return true
}

func validateOR(claimsConfig *middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	if utils.IsEmptyOrNil(claimsConfig.OR) {
		return true
	}
	if !utils.IsEmptyOrNil(claimsConfig.OR) {
		for _, v := range claimsConfig.OR {
			if v.Directive == middleware_oidc.ClaimTypeAndValue {
				if claimsPrincipal.HasClaim(v.Claim) {
					return true
				}
			}
			if v.Directive == middleware_oidc.ClaimType {
				if claimsPrincipal.HasClaimType(v.Claim.Type) {
					return true
				}
			}
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
func FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption(grpcEntrypointClaimsMap map[string]*middleware_oidc.EntryPointConfig, enableZeroTrust bool) grpc.UnaryServerInterceptor {

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
		subLogger = subLogger.With().Caller().Logger()

		subLogger.Debug().Msg("Enter")
		defer subLogger.Debug().Msg("Enter")

		if requestContainer != nil {
			claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)

			permissionDeniedFunc := func() (interface{}, error) {
				subLogger.Debug().Msg("Permission denied")
				return nil, status.Errorf(codes.PermissionDenied, "permission denied")
			}
			elem, ok := grpcEntrypointClaimsMap[info.FullMethod]
			if !ok && enableZeroTrust {
				subLogger.Debug().Msg("FullMethod not found in entrypoint claims map")
				return permissionDeniedFunc()
			}
			if !Validate(&subLogger, elem.ClaimsConfig, claimsPrincipal) {
				subLogger.Debug().Msg("ClaimsConfig validation failed")
				return permissionDeniedFunc()
			}
		}
		return handler(ctx, req)
	}
}

// FinalAuthVerificationMiddlewareUsingClaimsMap evaluates the claims principal
func FinalAuthVerificationMiddlewareUsingClaimsMap(grpcEntrypointClaimsMap map[string]*middleware_oidc.EntryPointConfig) grpc.UnaryServerInterceptor {
	return FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption(grpcEntrypointClaimsMap, false)
}

// FinalAuthVerificationMiddlewareUsingClaimsMapWithZeroTrust evaluates the claims principal
func FinalAuthVerificationMiddlewareUsingClaimsMapWithZeroTrust(grpcEntrypointClaimsMap map[string]*middleware_oidc.EntryPointConfig) grpc.UnaryServerInterceptor {
	return FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption(grpcEntrypointClaimsMap, true)
}
