package contextaccessor

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IContextAccessor,IInternalContextAccessor"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IContextAccessor,IInternalContextAccessor

import (
	"context"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// IContextAccessor ...
type IContextAccessor interface {
	GetContext() context.Context
	GetContainer() di.Container
}

// IInternalContextAccessor ...
type IInternalContextAccessor interface {
	IContextAccessor
	SetContext(context.Context)
}
