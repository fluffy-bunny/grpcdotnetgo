package serviceprovider

import (
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type IServiceProvider interface {
	GetService(name string) interface{}
}

type serviceProvider struct {
	Container di.Container
	Logger    servicesLogger.ILogger
}

func (s *serviceProvider) GetService(name string) interface{} {
	service, err := s.Container.SafeGet(name)
	if err != nil {
		s.Logger.Error().Err(err).Str("service-name", name).Msg("Failed to get service")
		return nil
	}
	return service
}
