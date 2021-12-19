package oauth2

import (
	loggerContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
)

type service struct {
	Request contracts_request.IRequest `inject:""`
	Logger  loggerContracts.ILogger    `inject:""`
}
