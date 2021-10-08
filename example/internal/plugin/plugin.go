package plugin

import (
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal/startup"
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
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
func (p *pluginService) GetStartup() coreContracts.IStartup {
	return startup.NewStartup()
}
