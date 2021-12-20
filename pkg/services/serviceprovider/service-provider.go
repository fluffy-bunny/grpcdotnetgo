package serviceprovider

import (
	"reflect"

	contracts_serviceprovider "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/serviceprovider"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type serviceScoped struct {
	container di.Container
}
type serviceSingleton struct {
	container di.Container
}

// AddScopedIServiceProvider adds service to the DI container
func AddScopedIServiceProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddScopedIServiceProvider")
	contracts_serviceprovider.AddScopedIServiceProvider(builder, reflect.TypeOf(&serviceScoped{}))
}

// AddSingletonISingletonServiceProvider adds service to the DI container
func AddSingletonISingletonServiceProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddSingletonIServiceProvider")
	contracts_serviceprovider.AddSingletonISingletonServiceProvider(builder, reflect.TypeOf(&serviceSingleton{}))
}

// AddServiceProviders adds service to the DI container
func AddServiceProviders(builder *di.Builder) {
	AddScopedIServiceProvider(builder)
	AddSingletonISingletonServiceProvider(builder)
}

func (s *serviceScoped) GetContainer() di.Container {
	return s.container
}
func (s *serviceScoped) SetContainer(container di.Container) {
	s.container = container
}
func (s *serviceSingleton) GetContainer() di.Container {
	return s.container
}
func (s *serviceSingleton) SetContainer(container di.Container) {
	s.container = container
}
