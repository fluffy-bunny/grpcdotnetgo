package plugin

import (
	"sync"

	pluginContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"
)

var (
	mutex            sync.Mutex
	availablePlugins []pluginContracts.IGRPCDotNetGoPlugin
)

// AddPlugin adds a plugin to the master registry
func AddPlugin(plugin pluginContracts.IGRPCDotNetGoPlugin) {
	mutex.Lock()
	defer mutex.Unlock()
	availablePlugins = append(availablePlugins, plugin)
	//log.Debug().Interface("availablePlugins", availablePlugins).Send()
}

// GetPlugins gets all the plugins
func GetPlugins() []pluginContracts.IGRPCDotNetGoPlugin {
	mutex.Lock()
	defer mutex.Unlock()
	return availablePlugins
}

// ClearPlugins usually used only for testing
func ClearPlugins() {
	mutex.Lock()
	defer mutex.Unlock()
	availablePlugins = make([]pluginContracts.IGRPCDotNetGoPlugin, 0)
}
