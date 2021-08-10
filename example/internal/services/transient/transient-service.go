package transient

import (
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
)

// Service is used to implement helloworld.GreeterServer.
type Service struct {
	ServiceProvider servicesServiceProvider.IServiceProvider
	name            string
	config          *internal.Config
}

// SetName ...
func (s *Service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *Service) GetName() string {
	return s.name
}

type Service2 struct {
	ServiceProvider servicesServiceProvider.IServiceProvider
	name            string
	config          *internal.Config
}

// SetName ...
func (s *Service2) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *Service2) GetName() string {
	return s.name
}
