package dicontext

import (
	"context"
	"time"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	contractsContextAccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/contextaccessor"
	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptors that performs request rate limiting.
func UnaryServerInterceptor(rootContainer di.Container) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Create a request and delete it once it has been handled.
		// Deleting the request will close the connection.
		requestContainer, _ := rootContainer.SubContainer()
		defer requestContainer.Delete()

		ctx = dicontext.SetRequestContainer(ctx, requestContainer)

		contextaccessor := contractsContextAccessor.GetIInternalContextAccessorFromContainer(requestContainer)
		contextaccessor.SetContext(ctx)

		// get a fresh ClaimsPrincipal from the request container and populate it with uuid data
		// this ensures that this claims principal object lives for the lifetime of the request
		claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)
		claimsPrincipal.AddClaim(claimsprincipalContracts.Claim{
			Type:  "_requestTime",
			Value: time.Now().UTC().String(),
		})
		return handler(ctx, req)
	}
}
