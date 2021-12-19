package dicontext

import (
	"context"
	"time"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
	contracts_serviceprovider "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/serviceprovider"
	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
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

		// expose the request container in an IServiceProvider object
		serviceProvider := contracts_serviceprovider.GetIServiceProviderFromContainer(requestContainer)
		serviceProviderInternal := serviceProvider.(contracts_serviceprovider.IServiceProviderInternal)
		serviceProviderInternal.SetContainer(requestContainer)

		request := contracts_request.GetIRequestFromContainer(requestContainer)
		innerRequest := request.(contracts_request.IInnerRequest)

		innerRequest.SetUnaryServerInfo(info)
		innerRequest.SetMetadata(metautils.ExtractIncoming(ctx))
		innerRequest.SetContext(ctx)

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
