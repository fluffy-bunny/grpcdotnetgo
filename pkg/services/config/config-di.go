package config

import (
	di "github.com/fluffy-bunny/sarulabsdi"
)

var (
	diName = di.GenerateUniqueServiceKey("grpcdotnetgo.config.")
)

// GetConfigFromContainer from the Container
func GetConfigFromContainer(ctn di.Container) interface{} {
	obj := ctn.Get(diName).(interface{})
	return obj
}

// AddConfig adds service to the DI container
func AddConfig(builder *di.Builder, config interface{}) {
	builder.Add(di.Def{
		Scope: di.App,
		Name:  diName,
		Build: func(ctn di.Container) (interface{}, error) {
			return config, nil
		},
	})
}
