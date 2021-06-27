package transient

import (
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
)

// Service is used to implement helloworld.GreeterServer.
type service struct {
	ServiceProvider servicesServiceProvider.IServiceProvider
	name            string
}

// SetName ...
func (s *service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service) GetName() string {
	return s.name
}

type service2 struct {
	ServiceProvider servicesServiceProvider.IServiceProvider
	name            string
}

// SetName ...
func (s *service2) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service2) GetName() string {
	return s.name
}
