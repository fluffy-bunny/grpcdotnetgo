package oauth2

import (
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
)

type IOauth2 interface {
}
type service struct {
	ContextAccessor contextaccessor.IContextAccessor
	Logger          servicesLogger.ILogger
}
