package transient

import (
	contracts_config "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/config"
	contracts_transient "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/transient"
)

// Service is used to implement helloworld.GreeterServer.
type (
	service struct {
		name   string
		config *contracts_config.Config
	}
	service2 struct {
		name   string
		config *contracts_config.Config
	}
)

func buildBreak() contracts_transient.ITransient {
	return &service{}
}
func buildBreak2() contracts_transient.ITransient {
	return &service2{}
}

// SetName ...
func (s *service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service) GetName() string {
	return "service-transient:" + s.name
}

// SetName ...
func (s *service2) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *service2) GetName() string {
	return "service2-transient:" + s.name
}
