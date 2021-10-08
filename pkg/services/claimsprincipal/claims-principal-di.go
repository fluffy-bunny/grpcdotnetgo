package claimsprincipal

import (
	"reflect"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

var (
	rtClaimsPrincipal           = reflect.TypeOf(&claimsPrincipal{}).Elem()
	reflectTypeIClaimsPrincipal = di.GetInterfaceReflectType((*claimsprincipalContracts.IClaimsPrincipal)(nil))
)

// GetClaimsPrincipalFromContainer from the Container
func GetClaimsPrincipalFromContainer(ctn di.Container) claimsprincipalContracts.IClaimsPrincipal {
	return ctn.GetByType(reflectTypeIClaimsPrincipal).(claimsprincipalContracts.IClaimsPrincipal)
}

// AddClaimsPrincipal adds service to the DI container
func AddClaimsPrincipal(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddClaimsPrincipal")
	implementedTypes := di.NewTypeSet()
	implementedTypes.Add(reflectTypeIClaimsPrincipal)
	builder.Add(di.Def{
		Type:             reflect.TypeOf(&claimsPrincipal{}),
		ImplementedTypes: implementedTypes,
		Scope:            di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &claimsPrincipal{
				claims: make(map[string][]string),
			}, nil
		},
	})
}
