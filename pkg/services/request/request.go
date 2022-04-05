package request

import (
	"context"
	"reflect"

	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
	contracts_serviceprovider "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/serviceprovider"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
)

type serviceRequest struct {
	Items           contracts_request.IItems                   `inject:""`
	ClaimsPrincipal contracts_claimsprincipal.IClaimsPrincipal `inject:""`
	ServiceProvider contracts_serviceprovider.IServiceProvider `inject:""`
	context         context.Context
	md              metautils.NiceMD
	unaryServerInfo *grpc.UnaryServerInfo
}

// AddScopedIRequest adds service to the DI container
func AddScopedIRequest(builder *di.Builder) {
	contracts_request.AddScopedIRequest(builder, reflect.TypeOf(&serviceRequest{}))
}
func (s *serviceRequest) GetUnaryServerInfo() *grpc.UnaryServerInfo {
	return s.unaryServerInfo
}
func (s *serviceRequest) SetUnaryServerInfo(info *grpc.UnaryServerInfo) {
	s.unaryServerInfo = info
}
func (s *serviceRequest) GetMetadata() metautils.NiceMD {
	return s.md
}

func (s *serviceRequest) GetItems() contracts_request.IItems {
	return nil
}
func (s *serviceRequest) SetMetadata(md metautils.NiceMD) {
	s.md = md
}
func (s *serviceRequest) GetContainer() di.Container {
	return dicontext.GetRequestContainer(s.context)
}
func (s *serviceRequest) GetContext() context.Context {
	return s.context
}
func (s *serviceRequest) SetContext(ctx context.Context) {
	s.context = ctx
}
func (s *serviceRequest) GetClaimsPrincipal() contracts_claimsprincipal.IClaimsPrincipal {
	return s.ClaimsPrincipal
}
func (s *serviceRequest) GetServiceProvider() contracts_serviceprovider.IServiceProvider {
	return s.ServiceProvider
}
