package dicontext

import (
	"context"

	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const ctxRequestContainer = "ctx-request-container"

func GetRequestContainer(ctx context.Context) di.Container {
	requestContainer := ctx.Value(ctxRequestContainer).(di.Container)
	return requestContainer
}
func setRequestContainer(ctx context.Context, requestContainer di.Container) context.Context {
	ctx = context.WithValue(ctx, ctxRequestContainer, requestContainer)
	return ctx
}

// UnaryServerInterceptor returns a new unary server interceptors that performs request rate limiting.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Create a request and delete it once it has been handled.
		// Deleting the request will close the connection.
		requestContainer, _ := grpcdotnetgo.GetContainer().SubContainer()
		defer requestContainer.Delete()

		ctx = setRequestContainer(ctx, requestContainer)

		contextaccessor := contextaccessor.GetInternalGetContextAccessorFromContainer(requestContainer)
		contextaccessor.SetContext(ctx)

		// get a fresh ClaimsPrincipal from the request container and populate it with uuid data

		claimsPrincipal := claimsprincipal.GetClaimsPrincipalFromContainer(requestContainer)
		claimsPrincipal.AddClaim(claimsprincipal.Claim{
			Type:  "d",
			Value: uuid.New().String(),
		})

		return handler(ctx, req)
	}
}
