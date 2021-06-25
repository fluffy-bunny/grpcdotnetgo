package claimsprincipal

import (
	"github.com/rs/zerolog/log"
	di "github.com/sarulabs/di/v2"
)

// Define an object in the App scope.

var diServiceName = "claims-principal"

// GetClaimsPrincipalFromContainer from the Container
func GetClaimsPrincipalFromContainer(ctn di.Container) *ClaimsPrincipal {
	return ctn.Get(diServiceName).(*ClaimsPrincipal)
}

// ClaimsPrincipal adds service to the DI container
func AddClaimsPrincipal(builder *di.Builder) {
	log.Info().Msg("IoC: ClaimsPrincipal")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &ClaimsPrincipal{}, nil
		},
	})
}
