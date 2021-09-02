package plugin

import (
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal/startup"
	grpcdotnetgo_core_types "github.com/fluffy-bunny/grpcdotnetgo/pkg/core/types"
	grpcdotnetgo_plugin "github.com/fluffy-bunny/grpcdotnetgo/pkg/plugin"
	grpcdotnetgo_plugin_types "github.com/fluffy-bunny/grpcdotnetgo/pkg/plugin/types"
)

func init() {
	grpcdotnetgo_plugin.AddPlugin(NewPlugin())
}

type pluginService struct {
}

func NewPlugin() grpcdotnetgo_plugin_types.IGRPCDotNetGoPlugin {
	return &pluginService{}
}
func (p *pluginService) GetName() string {
	return "example"
}
func (p *pluginService) GetStartup() grpcdotnetgo_core_types.IStartup {
	return startup.NewStartup()
}
