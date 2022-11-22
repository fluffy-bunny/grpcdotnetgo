package claimsprincipal

import (
	"context"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gogo/status"
	"github.com/rs/zerolog"
	zLog "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Validate ...
func Validate(ctx context.Context, claimsConfig *middleware_oidc.ClaimsConfig, claimsPrincipal claimsprincipalContracts.IClaimsPrincipal) bool {
	log := zerolog.Ctx(ctx).With().Logger()
	if !validateAND(claimsConfig, claimsPrincipal) {
		log.Debug().Msg("AND validation failed")
		return false
	}

	if !validateOR(claimsConfig, claimsPrincipal) {
		log.Debug().Msg("OR validation failed")
		return false
	}
	if claimsConfig.Child != nil {
		return Validate(ctx, claimsConfig.Child, claimsPrincipal)
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

	zLog.Info().Interface("entryPointConfig", grpcEntrypointClaimsMap).Send()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log := zerolog.Ctx(ctx).With().Bool("enableZeroTrust", enableZeroTrust).Logger()
		ctx = log.WithContext(ctx)
		requestContainer := middleware_dicontext.GetRequestContainer(ctx)
		log.Debug().Msg("FinalAuthVerificationMiddlewareUsingClaimsMapWithTrustOption")
		if requestContainer != nil {
			claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)

			permissionDeniedFunc := func() (interface{}, error) {
				log.Debug().Msg("Permission denied")
				return nil, status.Errorf(codes.PermissionDenied, "permission denied")
			}
			elem, ok := grpcEntrypointClaimsMap[info.FullMethod]
			if !ok && enableZeroTrust {
				log.Debug().Msg("FullMethod not found in entrypoint claims map")
				return permissionDeniedFunc()
			}
			if !Validate(ctx, elem.ClaimsConfig, claimsPrincipal) {
				log.Debug().Msg("ClaimsConfig validation failed")
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
