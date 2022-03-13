package scoped

import (
	contracts_config "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/config"
	contracts_scoped "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/scoped"
)

// Service is used to implement helloworld.GreeterServer.
type service struct {
	name   string
	config *contracts_config.Config
}

func buildBreak() contracts_scoped.IScoped {
	return &service{}
}

// SetName ...
func (s *service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service) GetName() string {
	return "service-scoped:" + s.name
}
