package mockoidcservice

import (
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"

	di "github.com/fluffy-bunny/sarulabsdi"

	"github.com/rs/zerolog/log"
)

var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("mockoidcservice")

func GetMockOIDCServiceFromContainer(ctn di.Container) interface{} {
	service := ctn.Get(diServiceName)
	return service
}

// AddMockOIDCService adds service to the DI container
func AddMockOIDCService(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceName).
		Msg("IoC: AddMockOIDCService")
	builder.Add(di.Def{
		Name:  diServiceName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			obj := &service{
				ServiceProvider: servicesServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}
			obj.Run()
			return obj, nil
		},
		Close: func(obj interface{}) error {
			log.Info().Msg("Closing Mock OIDC Service")
			service := obj.(*service)
			service.Shutdown()
			return nil
		},
	})
}
