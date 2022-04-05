package claimsprincipal

import (
	"reflect"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// AddScopedIClaimsPrincipal adds service to the DI container
func AddScopedIClaimsPrincipal(builder *di.Builder) {
	claimsprincipalContracts.AddScopedIClaimsPrincipal(builder,
		reflect.TypeOf(&claimsPrincipal{}))
}
