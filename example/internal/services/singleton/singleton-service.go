package singleton

import (
	contracts_config "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/config"
	contracts_singleton "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/singleton"
)

// Service is used to implement helloworld.GreeterServer.
type service struct {
	name   string
	config *contracts_config.Config
}

func buildBreak() contracts_singleton.ISingleton {
	return &service{}
}

// SetName ...
func (s *service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service) GetName() string {
	return "service-singleton:" + s.name
}
