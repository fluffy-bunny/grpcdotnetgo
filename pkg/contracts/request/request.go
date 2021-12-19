package request

import (
	"context"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
)

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IRequest,IItems"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IRequest,IItems

type (
	// IItems ...
	IItems interface {
		Set(key string, value interface{})
		Get(key string) interface{}
		Delete(key string)
		Clear()
		Keys() []string
		GetItems() map[string]interface{}
	}

	// IRequest ...
	IRequest interface {
		GetMetadata() metautils.NiceMD
		GetItems() IItems
		GetUnaryServerInfo() *grpc.UnaryServerInfo
		GetContext() context.Context
		GetContainer() di.Container
		GetClaimsPrincipal() contracts_claimsprincipal.IClaimsPrincipal
	}
	// IInnerRequest ...
	IInnerRequest interface {
		IRequest
		SetMetadata(md metautils.NiceMD)
		SetUnaryServerInfo(info *grpc.UnaryServerInfo)
		SetContext(context.Context)
	}
)
