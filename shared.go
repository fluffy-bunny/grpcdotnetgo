package grpcdotnetgo

import (
	claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	singletonServicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/singleton-serviceprovider"
	"github.com/sarulabs/di"
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
	builder, err := di.NewBuilder(di.App, di.Request)
	if err != nil {
		return nil, err
	}

	claimsprincipal.AddClaimsPrincipal(builder)
	contextaccessor.AddContextAccessor(builder)
	servicesLogger.AddRequestLogger(builder)
	servicesServiceProvider.AddServiceProvider(builder)
	singletonServicesServiceProvider.AddSingletonServiceProvider(builder)

	return &DotNetGoBuilder{
		Builder: builder,
	}, nil
}

func (b *DotNetGoBuilder) Build() {
	container = b.Builder.Build()
}
