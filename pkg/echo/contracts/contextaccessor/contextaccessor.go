package contextaccessor

import (
	"github.com/labstack/echo/v4"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IInternalEchoContextAccessor,IEchoContextAccessor"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/$GOPACKAGE IInternalEchoContextAccessor,IEchoContextAccessor

type (
	// IEchoContextAccessor ...
	IEchoContextAccessor interface {
		GetContext() echo.Context
	}
	// IInternalEchoContextAccessor ...
	IInternalEchoContextAccessor interface {
		IEchoContextAccessor
		SetContext(echo.Context)
	}
)
