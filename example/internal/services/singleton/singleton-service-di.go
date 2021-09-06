package singleton

import (
	"reflect"

	servicesConfig "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/config"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

var (
	rtGetType = reflect.TypeOf(&service{}).Elem()
)

// Define an object in the App scope.

// GetSingletonServiceFromContainer from the Container
func GetSingletonServiceFromContainer(ctn di.Container) *service {
	return ctn.GetByType(rtGetType).(*service)
}

// GreeterAddSingletonServiceService adds service to the DI container
func AddSingletonService(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddSingletonService")
	builder.Add(di.Def{
		Type:  reflect.TypeOf(&service{}),
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &service{
				config: servicesConfig.GetConfigFromContainer(ctn),
			}, nil
		},
	})
}
