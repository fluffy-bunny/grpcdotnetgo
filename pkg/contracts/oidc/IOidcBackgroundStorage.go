package oidc

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IOidcBackgroundStorage"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IOidcBackgroundStorage

import (
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
)

type (
	// IOidcBackgroundStorage ...
	IOidcBackgroundStorage interface {
		AtomicStore(disco *middleware_oidc.DiscoveryDocument)
		AtomicGet() *middleware_oidc.DiscoveryDocument
	}
)
