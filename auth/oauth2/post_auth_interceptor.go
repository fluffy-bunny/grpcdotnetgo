package oauth2

import (
	"context"
	"fmt"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/gogo/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func FinalAuthVerificationMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println(info.FullMethod)
		data := ctx.Value(CtxClaimsPrincipalKey)
		if data == nil {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}

		claimsPrincipal := data.(*ClaimsPrincipal)

		var permissions []string
		From(claimsPrincipal.Claims).Where(func(c interface{}) bool {
			return c.(Claim).Type == "permissions"
		}).Select(func(c interface{}) interface{} {
			return c.(Claim).Value
		}).ToSlice(&permissions)
		var mapPermissions = make(map[string]bool)
		for _, element := range permissions {
			mapPermissions[element] = true
		}
		newCtx := context.WithValue(ctx, CtxClaimsPermissions, mapPermissions)
		return newCtx, nil
	}

	//return grpc_auth.UnaryServerInterceptor(buildFinalAuthVerificationFunc())
}

func buildFinalAuthVerificationFunc() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		data := ctx.Value(CtxClaimsPrincipalKey)
		if data == nil {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}

		claimsPrincipal := data.(*ClaimsPrincipal)

		var permissions []string
		From(claimsPrincipal.Claims).Where(func(c interface{}) bool {
			return c.(Claim).Type == "permissions"
		}).Select(func(c interface{}) interface{} {
			return c.(Claim).Value
		}).ToSlice(&permissions)
		var mapPermissions = make(map[string]bool)
		for _, element := range permissions {
			mapPermissions[element] = true
		}
		newCtx := context.WithValue(ctx, CtxClaimsPermissions, mapPermissions)
		return newCtx, nil

	}
}
