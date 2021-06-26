package singleton

import (
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
)

// Service is used to implement helloworld.GreeterServer.
type Service struct {
	ServiceProvider servicesServiceProvider.IServiceProvider
	name            string
}

// SetName ...
func (s *Service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *Service) GetName() string {
	return s.name
}
