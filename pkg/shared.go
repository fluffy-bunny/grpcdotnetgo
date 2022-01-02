package pkg

import (
	services_BackgroundTasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	services_Logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	services_metadatafilter "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/metadatafilter"
	services_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/request"
	di "github.com/fluffy-bunny/sarulabsdi"
	_ "github.com/gogo/protobuf/gogoproto" // ensures that go mod vendor brings it down
)

const (
	// SupportPackageIsVersion7 ...
	SupportPackageIsVersion7 = true
)

// DotNetGoBuilder ...
type DotNetGoBuilder struct {
	Builder   *di.Builder
	Container di.Container
}

// NewDotNetGoBuilder ...
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
	services_claimsprincipal.AddScopedIClaimsPrincipal(builder)
	services_request.AddScopedIRequest(builder)
	services_request.AddScopedIItems(builder)
	services_Logger.AddScopedILogger(builder)
	services_Logger.AddSingletonILogger(builder)
	services_BackgroundTasks.AddSingletonBackgroundTasks(builder)
	services_metadatafilter.AddSingletonIMetadataFilterMiddlewareNil(builder)
}

// Build ...
func (dngbuilder *DotNetGoBuilder) Build() di.Container {
	dngbuilder.Container = dngbuilder.Builder.Build()
	return dngbuilder.Container
}
