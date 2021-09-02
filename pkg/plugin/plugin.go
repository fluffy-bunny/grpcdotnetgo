package plugin

import (
	grpcdotnetgo_plugin_types "github.com/fluffy-bunny/grpcdotnetgo/pkg/plugin/types"
	"github.com/rs/zerolog/log"
)

var (
	availablePlugins []grpcdotnetgo_plugin_types.IGRPCDotNetGoPlugin
)

func AddPlugin(plugin grpcdotnetgo_plugin_types.IGRPCDotNetGoPlugin) {
	availablePlugins = append(availablePlugins, plugin)
	log.Debug().Interface("availablePlugins", availablePlugins).Send()
}
func GetPlugins() []grpcdotnetgo_plugin_types.IGRPCDotNetGoPlugin {
	return availablePlugins
}
