package oauth2

import (
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
)

// IOauth2 contract
type IOauth2 interface {
}
type service struct {
	ContextAccessor contextaccessor.IContextAccessor
	Logger          loggerContracts.ILogger
}
