package container

import (
	di "github.com/fluffy-bunny/sarulabsdi"
)

//go:generate genny   -pkg $GOPACKAGE     -in=../../../../genny/sarulabsdi/func-types.go -out=gen-func-$GOFILE gen "FuncType=ContainerAccessor"

type (
	ContainerAccessor func() di.Container
)
