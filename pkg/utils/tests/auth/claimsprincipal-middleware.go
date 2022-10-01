package auth

import (
	"context"
	"encoding/json"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	core_jwt "github.com/fluffy-bunny/grpcdotnetgo/pkg/core/jwt"
	grpc_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/go-grpc-middleware/auth"

	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// EntryPointToClaimsMap configuration
type EntryPointToClaimsMap map[string][]contracts_claimsprincipal.Claim

// TestClaimsPrincipalProducerUnaryServerInterceptor ...
func TestClaimsPrincipalProducerUnaryServerInterceptor(claimsMap EntryPointToClaimsMap) grpc.UnaryServerInterceptor {
	authFunc := buildTestAuthFunction(claimsMap)
	return grpc_auth.UnaryServerInterceptor(authFunc)
}
func buildTestAuthFunction(claimsMap EntryPointToClaimsMap) func(ctx context.Context, fullMethodName string) (context.Context, error) {
	return func(ctx context.Context, fullMethodName string) (context.Context, error) {
		requestContainer := dicontext.GetRequestContainer(ctx)
		claimsPrincipal := contracts_claimsprincipal.GetIClaimsPrincipalFromContainer(requestContainer)
		claims, ok := claimsMap[fullMethodName]
		if ok {
			for _, c := range claims {
				claimsPrincipal.AddClaim(c)
			}
		}
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			unitTestClaims := md.Get("x-unit-test-claims")
			if !utils.IsEmptyOrNil(unitTestClaims) {
				unitTestClaims2 := []contracts_claimsprincipal.Claim{}
				err := json.Unmarshal([]byte(unitTestClaims[0]), &unitTestClaims2)
				if err == nil {
					for _, c := range unitTestClaims2 {
						claimsPrincipal.AddClaim(c)
					}
				}
			}
		}
		return ctx, nil
	}
}

// TestClaimsPrincipalUnsignedJWTServerInterceptor ...
func TestClaimsPrincipalUnsignedJWTServerInterceptor() grpc.UnaryServerInterceptor {
	authFunc := buildUnsignedJWTAuthFunction()
	return grpc_auth.UnaryServerInterceptor(authFunc)
}
func buildUnsignedJWTAuthFunction() func(ctx context.Context, fullMethodName string) (context.Context, error) {
	return func(ctx context.Context, fullMethodName string) (context.Context, error) {
		requestContainer := dicontext.GetRequestContainer(ctx)
		claimsPrincipal := contracts_claimsprincipal.GetIClaimsPrincipalFromContainer(requestContainer)

		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil || utils.IsEmptyOrNil(token) {
			// nothing here
			return ctx, nil
		}

		claimsPrincipal2, err := core_jwt.ClaimsPrincipalFromUnsignedToken(token)
		if err != nil {
			// nothing here
			return ctx, nil
		}
		claimsPrincipal.AddClaim(claimsPrincipal2.GetClaims()...)

		return ctx, nil
	}
}
