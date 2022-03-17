package inmemory

import (
	"reflect"

	contracts_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/auth"
	coreUtilsTestsAuth "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils/tests/auth"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

type (
	service struct {
		claimsmap coreUtilsTestsAuth.EntryPointToClaimsMap
	}
)

func buildBreak() contracts_auth.IModularAuthMiddleware {
	return &service{}
}

// AddSingletonIModularAuthMiddleware ...
func AddSingletonIModularAuthMiddleware(builder *di.Builder, claimsmap coreUtilsTestsAuth.EntryPointToClaimsMap) {
	contracts_auth.AddSingletonIModularAuthMiddlewareByFunc(builder,
		reflect.TypeOf(&service{}),
		func(ctn di.Container) (interface{}, error) {
			return &service{
				claimsmap: claimsmap,
			}, nil
		})
}

func (s *service) GetUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return coreUtilsTestsAuth.TestClaimsPrincipalProducerUnaryServerInterceptor(s.claimsmap)
}
