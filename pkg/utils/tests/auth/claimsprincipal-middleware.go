package auth

import (
	"context"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	grpc_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/go-grpc-middleware/auth"
	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	"google.golang.org/grpc"
)

// EntryPointToClaimsMap configuration
type EntryPointToClaimsMap map[string][]claimsprincipalContracts.Claim

// TestClaimsPrincipalProducerUnaryServerInterceptor ...
func TestClaimsPrincipalProducerUnaryServerInterceptor(claimsMap EntryPointToClaimsMap) grpc.UnaryServerInterceptor {
	authFunc := buildTestAuthFunction(claimsMap)
	return grpc_auth.UnaryServerInterceptor(authFunc)
}
func buildTestAuthFunction(claimsMap EntryPointToClaimsMap) func(ctx context.Context, fullMethodName string) (context.Context, error) {
	return func(ctx context.Context, fullMethodName string) (context.Context, error) {
		requestContainer := dicontext.GetRequestContainer(ctx)
		claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)
		claims, ok := claimsMap[fullMethodName]
		if ok {
			for _, c := range claims {
				claimsPrincipal.AddClaim(c)
			}
		}
		return ctx, nil
	}
}
