package inmemory

import (
	"reflect"

	contracts_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/auth"
	coreUtilsTestsAuth "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils/tests/auth"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type (
	service struct {
		claimsmap coreUtilsTestsAuth.EntryPointToClaimsMap
	}
)

func assertImplementation() {
	var _ contracts_auth.IModularAuthMiddleware = (*service)(nil)
}

// AddSingletonIModularAuthMiddleware ...
func AddSingletonIModularAuthMiddleware(builder *di.Builder, claimsmap coreUtilsTestsAuth.EntryPointToClaimsMap) {
	log.Info().Msg("IoC: AddSingletonIModularAuthMiddleware")
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
