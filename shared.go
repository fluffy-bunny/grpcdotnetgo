package grpcdotnetgo

import (
	claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	di "github.com/fluffy-bunny/sarulabsdi"
)

const (
	SupportPackageIsVersion7 = true
)

// Container for our IoC
var container di.Container

type DotNetGoBuilder struct {
	Builder *di.Builder
}

func GetContainer() di.Container {
	return container
}

func NewDotNetGoBuilder() (*DotNetGoBuilder, error) {
	builder, err := di.NewBuilder(di.App, di.Request, "transient")
	if err != nil {
		return nil, err
	}

	claimsprincipal.AddClaimsPrincipal(builder)
	contextaccessor.AddContextAccessor(builder)
	servicesLogger.AddScopedLogger(builder)
	servicesLogger.AddSingletonLogger(builder)
	servicesServiceProvider.AddScopedServiceProvider(builder)
	servicesServiceProvider.AddSingletonServiceProvider(builder)

	return &DotNetGoBuilder{
		Builder: builder,
	}, nil
}

func (b *DotNetGoBuilder) Build() di.Container {
	container = b.Builder.Build()
	return container
}
