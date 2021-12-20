package serviceprovider

import di "github.com/fluffy-bunny/sarulabsdi"

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IServiceProvider,ISingletonServiceProvider"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IServiceProvider,ISingletonServiceProvider

type (
	// IServiceProvider interface
	IServiceProvider interface {
		GetContainer() di.Container
	}
	// IServiceProviderInternal interface
	IServiceProviderInternal interface {
		IServiceProvider
		SetContainer(di.Container)
	}
	// ISingletonServiceProvider interface
	ISingletonServiceProvider interface {
		GetContainer() di.Container
	}
	// ISingletonServiceProviderInternal interface
	ISingletonServiceProviderInternal interface {
		ISingletonServiceProvider
		SetContainer(di.Container)
	}
)
