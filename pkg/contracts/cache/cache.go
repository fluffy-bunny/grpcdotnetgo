package cache

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ICache,IMemoryCache"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ICache,IMemoryCache

import (
	gookit_cache "github.com/gookit/cache"
)

type (
	// ICache interface
	ICache interface {
		gookit_cache.Cache
	}
	// IMemoryCache interface
	IMemoryCache interface {
		ICache
	}
)
