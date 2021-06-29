package transient

import (
	"reflect"

	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	exampleServices "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// GetTransientServiceFromContainer from the Container
func GetTransientService() *Service {
	return GetTransientServiceFromContainer(grpcdotnetgo.GetContainer())
}

// GetTransientServiceFromContainer from the Container
func GetTransientServiceFromContainer(ctn di.Container) *Service {
	return ctn.GetByType(reflect.TypeOf(&Service{}).Elem()).(*Service)
}

// GreeterAddTransientServiceService adds service to the DI container
func AddTransientService(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddTransientService")

	types := di.NewTypeSet()
	inter := di.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	types.Add(reflect.TypeOf(&Service{}))

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service{}),
		Unshared:         true,
		Build: func(ctn di.Container) (interface{}, error) {
			service := &Service{
				ServiceProvider: servicesServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}
			return service, nil
		},
	})
}

// GetTransientServiceFromContainer from the Container
func GetTransientService2() *Service2 {
	return GetTransientService2FromContainer(grpcdotnetgo.GetContainer())
}

// GetTransientServiceFromContainer from the Container
func GetTransientService2FromContainer(ctn di.Container) *Service2 {
	return ctn.GetByType(reflect.TypeOf(&Service2{}).Elem()).(*Service2)
}

func AddTransientService2(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddTransientService2")

	types := di.NewTypeSet()
	inter := di.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	types.Add(reflect.TypeOf(&Service2{}))

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service2{}),
		Unshared:         true,
		Build: func(ctn di.Container) (interface{}, error) {
			service := &Service2{
				ServiceProvider: servicesServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}
			return service, nil
		},
	})
}
