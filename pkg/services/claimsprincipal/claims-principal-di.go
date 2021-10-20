package claimsprincipal

import (
	"reflect"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddClaimsPrincipal adds service to the DI container
func AddClaimsPrincipal(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddClaimsPrincipal")
	claimsprincipalContracts.AddSingletonIClaimsPrincipalByFunc(builder,
		reflect.TypeOf(&claimsPrincipal{}), func(ctn di.Container) (interface{}, error) {
			return &claimsPrincipal{
				claims: make(map[string][]string),
			}, nil
		})
}
