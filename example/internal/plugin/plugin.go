package plugin

import (
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal/startup"
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	pluginContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"
	grpcdotnetgo_plugin "github.com/fluffy-bunny/grpcdotnetgo/pkg/plugin"
)

func init() {
	grpcdotnetgo_plugin.AddPlugin(NewPlugin())
}

type pluginService struct {
}

// NewPlugin creates a new plugin
func NewPlugin() pluginContracts.IGRPCDotNetGoPlugin {
	return &pluginService{}
}

// GetName gets name of the plugin
func (p *pluginService) GetName() string {
	return "example"
}

// GetStartup gets the plugin's IStartup
func (p *pluginService) GetStartup() coreContracts.IStartup {
	return startup.NewStartup()
}
