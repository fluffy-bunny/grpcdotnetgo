package grpc

import (
	"google.golang.org/grpc"
)

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IServiceEndpointRegistration"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IServiceEndpointRegistration

type (
	// IServiceEndpointRegistration interface
	IServiceEndpointRegistration interface {
		GetName() string
		GetNewClient(cc grpc.ClientConnInterface) interface{}
		RegisterEndpoint(server *grpc.Server) interface{}
	}
)
