package transient

import (
	"reflect"

	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	exampleServices "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services"
	grpcdotnetgreflect "github.com/fluffy-bunny/grpcdotnetgo/reflect"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("di-transient-service")

// GetTransientServiceFromContainer from the Container
func GetTransientService() *service {
	return GetTransientServiceFromContainer(grpcdotnetgo.GetContainer())
}

// GetTransientServiceFromContainer from the Container
func GetTransientServiceFromContainer(ctn di.Container) *service {
	return ctn.Get(diServiceName).(*service)
}

// GreeterAddTransientServiceService adds service to the DI container
func AddTransientService(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceName).
		Msg("IoC: AddTransientService")

	types := di.NewTypeSet()
	inter := grpcdotnetgreflect.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	types.Add(reflect.TypeOf(&service{}))

	builder.Add(di.Def{
		Name:             diServiceName,
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&service{}),
		Build: func(ctn di.Container) (interface{}, error) {
			service := &service{
				ServiceProvider: servicesServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}
			return service, nil
		},
	})
}

// Define an object in the App scope.
var diServiceName2 = grpcdotnetgoutils.GenerateUnqueServiceName("di-transient-service2")

// GetTransientServiceFromContainer from the Container
func GetTransientService2() *service2 {
	return GetTransientService2FromContainer(grpcdotnetgo.GetContainer())
}

// GetTransientServiceFromContainer from the Container
func GetTransientService2FromContainer(ctn di.Container) *service2 {
	return ctn.Get(diServiceName2).(*service2)
}

func AddTransientService2(builder *di.Builder) {
	log.Info().
		Str("serviceName", diServiceName2).
		Msg("IoC: AddTransientService2")

	types := di.NewTypeSet()
	inter := grpcdotnetgreflect.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	types.Add(reflect.TypeOf(&service2{}))

	builder.Add(di.Def{
		Name:             diServiceName2,
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&service2{}),
		Build: func(ctn di.Container) (interface{}, error) {
			service := &service2{
				ServiceProvider: servicesServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}
			return service, nil
		},
	})
}
