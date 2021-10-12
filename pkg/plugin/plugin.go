package plugin

import (
	pluginContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"
	"github.com/rs/zerolog/log"
)

var (
	availablePlugins []pluginContracts.IGRPCDotNetGoPlugin
)

// AddPlugin adds a plugin to the master registry
func AddPlugin(plugin pluginContracts.IGRPCDotNetGoPlugin) {
	availablePlugins = append(availablePlugins, plugin)
	log.Debug().Interface("availablePlugins", availablePlugins).Send()
}

// GetPlugins gets all the plugins
func GetPlugins() []pluginContracts.IGRPCDotNetGoPlugin {
	return availablePlugins
}
