package singleton

import (
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
)

// Service is used to implement helloworld.GreeterServer.
type service struct {
	name   string
	config *internal.Config
}

// SetName ...
func (s *service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service) GetName() string {
	return s.name
}
