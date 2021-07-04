package cmd

import (
	"github.com/fluffy-bunny/grpcdotnetgo/cobracore/cmd/serve"
	"github.com/fluffy-bunny/grpcdotnetgo/core"
)

func Start(startup core.IStartup) {
	serve.SetStartup(startup)
	Execute()
}
