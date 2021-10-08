package pkg

import (
	servicesBackgroundTasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	_ "github.com/gogo/protobuf/gogoproto" // ensures that go mod vendor brings it down
)

const (
	SupportPackageIsVersion7 = true
)

type DotNetGoBuilder struct {
	Builder   *di.Builder
	Container di.Container
}

func NewDotNetGoBuilder() (*DotNetGoBuilder, error) {
	builder, err := di.NewBuilder(di.App, di.Request, "transient")
	if err != nil {
		return nil, err
	}
	return &DotNetGoBuilder{
		Builder: builder,
	}, nil

}

func (dngbuilder *DotNetGoBuilder) AddDefaultService() {
	builder := dngbuilder.Builder

	claimsprincipal.AddClaimsPrincipal(builder)
	contextaccessor.AddContextAccessor(builder)
	servicesLogger.AddScopedLogger(builder)
	servicesLogger.AddSingletonLogger(builder)
	servicesBackgroundTasks.AddBackgroundTasks(builder)
}

func (b *DotNetGoBuilder) Build() di.Container {
	b.Container = b.Builder.Build()
	return b.Container
}
