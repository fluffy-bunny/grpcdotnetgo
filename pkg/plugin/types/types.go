package types

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
)

// IGRPCDotNetGoPlugin ...
type IGRPCDotNetGoPlugin interface {
	GetName() string
	GetStartup() coreContracts.IStartup
}
