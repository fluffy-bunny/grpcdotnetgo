package claimsprincipal

import (
	"reflect"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddScopedIClaimsPrincipal adds service to the DI container
func AddScopedIClaimsPrincipal(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddClaimsPrincipal")
	claimsprincipalContracts.AddScopedIClaimsPrincipal(builder,
		reflect.TypeOf(&claimsPrincipal{}))
}
