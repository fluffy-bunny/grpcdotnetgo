package transient

import (
	"fmt"
	"reflect"

	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo"
	exampleServices "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services"
	grpcdotnetgreflect "github.com/fluffy-bunny/grpcdotnetgo/reflect"
	singletonServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/singleton-serviceprovider"
	grpcdotnetgoutils "github.com/fluffy-bunny/grpcdotnetgo/utils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// Define an object in the App scope.
var diServiceName = grpcdotnetgoutils.GenerateUnqueServiceName("di-transient-service")

// GetTransientServiceFromContainer from the Container
func GetTransientService() *Service {
	return GetTransientServiceFromContainer(grpcdotnetgo.GetContainer())
}

// GetTransientServiceFromContainer from the Container
func GetTransientServiceFromContainer(ctn di.Container) *Service {
	return ctn.Get(diServiceName).(*Service)
}

// GreeterAddTransientServiceService adds service to the DI container
func AddTransientService(builder *di.Builder) {
	log.Info().Msg("IoC: AddTransientService")

	types := di.NewTypeSet()
	inter := grpcdotnetgreflect.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	types.Add(reflect.TypeOf(&Service{}))

	implType := reflect.TypeOf(&Service{})
	for rt := range types {
		if rt.Kind() == reflect.Interface {
			if !implType.Implements(rt) {
				panic(fmt.Errorf("%v does not implement %v", diServiceName, rt))
			}
		} else {
			if implType != rt {
				panic(fmt.Errorf("%v does not implement %v", diServiceName, rt))
			}
		}

	}

	builder.Add(di.Def{
		Name:             diServiceName,
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service{}),
		Build: func(ctn di.Container) (interface{}, error) {
			service := &Service{
				ServiceProvider: singletonServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}
			return service, nil
		},
	})
}

// Define an object in the App scope.
var diServiceName2 = grpcdotnetgoutils.GenerateUnqueServiceName("di-transient-service2")

// GetTransientServiceFromContainer from the Container
func GetTransientService2() *Service {
	return GetTransientServiceFromContainer(grpcdotnetgo.GetContainer())
}

// GetTransientServiceFromContainer from the Container
func GetTransientService2FromContainer(ctn di.Container) *Service {
	return ctn.Get(diServiceName2).(*Service)
}

func AddTransientService2(builder *di.Builder) {
	log.Info().Msg("IoC: AddTransientService")

	types := di.NewTypeSet()
	inter := grpcdotnetgreflect.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	types.Add(reflect.TypeOf(&Service2{}))

	builder.Add(di.Def{
		Name:             diServiceName2,
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service2{}),
		Build: func(ctn di.Container) (interface{}, error) {
			service := &Service2{
				ServiceProvider: singletonServiceProvider.GetSingletonServiceProviderFromContainer(ctn),
			}
			return service, nil
		},
	})
}
