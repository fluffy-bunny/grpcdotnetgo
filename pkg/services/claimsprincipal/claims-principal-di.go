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
	claimsprincipalContracts.AddScopedIClaimsPrincipal(builder, reflect.TypeOf(&claimsPrincipal{}))
}
