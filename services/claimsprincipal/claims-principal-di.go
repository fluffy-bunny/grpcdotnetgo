package claimsprincipal

import (
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("IClaimsPrincipal")

// GetClaimsPrincipalFromContainer from the Container
func GetClaimsPrincipalFromContainer(ctn di.Container) IClaimsPrincipal {
	return ctn.Get(diServiceName).(IClaimsPrincipal)
}

// ClaimsPrincipal adds service to the DI container
func AddClaimsPrincipal(builder *di.Builder) {
	log.Info().Msg("IoC: ClaimsPrincipal")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &claimsPrincipal{
				claims: make(map[string][]string),
			}, nil
		},
	})
}
