package serviceprovider

import (
	"fmt"
	"reflect"

	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type IServiceProvider interface {
	GetByName(name string) interface{}
	GetByType(rt reflect.Type) interface{}
	GetManyByType(rt reflect.Type) []interface{}
	GetContainer() di.Container
}

// IServiceProvider reflect type
var TypeIServiceProvider = di.GetInterfaceReflectType((*IServiceProvider)(nil))

type serviceProvider struct {
	Container di.Container
	Logger    servicesLogger.ILogger
}

func (s *serviceProvider) GetContainer() di.Container {
	return s.Container
}
func (s *serviceProvider) GetByType(rt reflect.Type) interface{} {
	service, err := s.Container.SafeGetByType(rt)
	if err != nil {
		s.Logger.Error().Err(err).Str("type", fmt.Sprintf("%v", rt)).Msg("Failed to get service")
		return nil
	}
	return service
}
func (s *serviceProvider) GetManyByType(rt reflect.Type) []interface{} {
	services, err := s.Container.SafeGetManyByType(rt)
	if err != nil {
		s.Logger.Error().Err(err).Str("type", fmt.Sprintf("%v", rt)).Msg("Failed to get service")
		return nil
	}
	return services
}
func (s *serviceProvider) GetByName(name string) interface{} {
	service, err := s.Container.SafeGet(name)
	if err != nil {
		s.Logger.Error().Err(err).Str("service-name", name).Msg("Failed to get service")
		return nil
	}
	return service
}
