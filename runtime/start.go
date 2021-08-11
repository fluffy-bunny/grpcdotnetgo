package runtime

import (
	"github.com/fluffy-bunny/grpcdotnetgo/cobracore/cmd"
	cmdVersion "github.com/fluffy-bunny/grpcdotnetgo/cobracore/cmd/version"
)

var Version string

func SetVersion(version string) {
	cmdVersion.SetVersion(version)
}
func Start() {
	cmd.Execute()
}
