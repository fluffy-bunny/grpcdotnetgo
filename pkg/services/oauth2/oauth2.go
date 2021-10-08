package oauth2

import (
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
)

type IOauth2 interface {
}
type service struct {
	ContextAccessor contextaccessor.IContextAccessor
	Logger          loggerContracts.ILogger
}
