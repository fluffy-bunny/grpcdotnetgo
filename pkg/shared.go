package pkg

import (
	contracts_serviceprovider "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/serviceprovider"
	services_BackgroundTasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/backgroundtasks"
	services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
	services_Logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	services_metadatafilter "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/metadatafilter"
	services_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/request"
	services_serviceprovider "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/serviceprovider"
	services_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/timeutils"
	services_uuid "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/uuid"
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
	services_serviceprovider.AddServiceProviders(builder)
	services_timeutils.AddSingletonITime(builder)
	services_timeutils.AddTimeNow(builder)
	services_uuid.AddSingletonIKSUID(builder)
}

// Build ...
func (dngbuilder *DotNetGoBuilder) Build() di.Container {
	dngbuilder.Container = dngbuilder.Builder.Build()
	serviceProvider := contracts_serviceprovider.GetISingletonServiceProviderFromContainer(dngbuilder.Container)
	serviceProviderInternal := serviceProvider.(contracts_serviceprovider.ISingletonServiceProviderInternal)
	serviceProviderInternal.SetContainer(dngbuilder.Container)

	return dngbuilder.Container
}
