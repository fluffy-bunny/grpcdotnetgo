package oauth2

import (
	contractsContextAccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/contextaccessor"
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
)

type service struct {
	ContextAccessor contractsContextAccessor.IContextAccessor `inject:""`
	Logger          loggerContracts.ILogger                   `inject:""`
}
