package dicontext

import (
	"context"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/google/uuid"
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

		contextaccessor := contextaccessor.GetInternalGetContextAccessorFromContainer(requestContainer)
		contextaccessor.SetContext(ctx)

		// get a fresh ClaimsPrincipal from the request container and populate it with uuid data

		claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)
		claimsPrincipal.AddClaim(claimsprincipalContracts.Claim{
			Type:  "d",
			Value: uuid.New().String(),
		})

		return handler(ctx, req)
	}
}
