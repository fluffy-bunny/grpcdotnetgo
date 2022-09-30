package auth

import (
	"context"
	"encoding/json"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
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
