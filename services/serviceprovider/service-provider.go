package serviceprovider

import (
	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
)

type IServiceProvider interface {
	GetService(name string) interface{}
}

type serviceProvider struct {
	Container di.Container
	Logger    *zerolog.Logger
}

func (s *serviceProvider) GetService(name string) interface{} {
	service, err := s.Container.SafeGet(name)
	if err != nil {
		s.Logger.Error().Err(err).Str("service-name", name).Msg("Failed to get service")
		return nil
	}
	return service
}
