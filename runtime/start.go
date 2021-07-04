package runtime

import (
	"github.com/fluffy-bunny/grpcdotnetgo/cobracore/cmd"
	"github.com/fluffy-bunny/grpcdotnetgo/cobracore/cmd/serve"
	cmdVersion "github.com/fluffy-bunny/grpcdotnetgo/cobracore/cmd/version"
	"github.com/fluffy-bunny/grpcdotnetgo/core"
)

var Version string

func SetVersion(version string) {
	cmdVersion.SetVersion(version)
}
func Start(startup core.IStartup) {
	serve.SetStartup(startup)
	cmd.Execute()
}
