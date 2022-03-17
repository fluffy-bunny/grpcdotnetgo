package auth

import (
	"google.golang.org/grpc"
)

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IModularAuthMiddleware"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE  github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IModularAuthMiddleware

type (
	// IModularAuthMiddleware ...
	IModularAuthMiddleware interface {
		// GetUnaryServerInterceptor ...
		GetUnaryServerInterceptor() grpc.UnaryServerInterceptor
	}
)
