package types

import (
	grpcdotnetgo_core_types "github.com/fluffy-bunny/grpcdotnetgo/pkg/core/types"
)

type IGRPCDotNetGoPlugin interface {
	GetName() string
	GetStartup() grpcdotnetgo_core_types.IStartup
}
