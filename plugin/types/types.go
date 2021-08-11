package types

import (
	grpcdotnetgo_core_types "github.com/fluffy-bunny/grpcdotnetgo/core/types"
)

type IGRPCDotNetGoPlugin interface {
	GetName() string
	GetStartup() grpcdotnetgo_core_types.IStartup
}
