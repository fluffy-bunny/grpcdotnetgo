package singleton

import (
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
)

// Service is used to implement helloworld.GreeterServer.
type service struct {
	ServiceProvider servicesServiceProvider.IServiceProvider
	name            string
	config          *internal.Config
}

// SetName ...
func (s *service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service) GetName() string {
	return s.name
}
