package pkg

import (
	servicesBackgroundTasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/contextaccessor"
	servicesMetadatafilter "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/metadatafilter"

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

// AddDefaultService ...
func (dngbuilder *DotNetGoBuilder) AddDefaultService() {
	builder := dngbuilder.Builder
	claimsprincipal.AddScopedIClaimsPrincipal(builder)
	contextaccessor.AddScopedContextAccessor(builder)
	servicesLogger.AddScopedILogger(builder)
	servicesLogger.AddSingletonILogger(builder)
	servicesBackgroundTasks.AddSingletonBackgroundTasks(builder)
	servicesMetadatafilter.AddSingletonIMetadataFilterMiddlewareNil(builder)
}

func (b *DotNetGoBuilder) Build() di.Container {
	b.Container = b.Builder.Build()
	return b.Container
}
