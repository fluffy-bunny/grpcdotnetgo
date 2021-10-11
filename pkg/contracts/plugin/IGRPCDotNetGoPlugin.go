package plugin

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
)

// IGRPCDotNetGoPlugin contract
type IGRPCDotNetGoPlugin interface {
	GetName() string
	GetStartup() coreContracts.IStartup
}
