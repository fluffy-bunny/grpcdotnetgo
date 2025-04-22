package health

import (
	"context"
	"reflect"

	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	di "github.com/fluffy-bunny/sarulabsdi"
	grpchealth "google.golang.org/grpc/health/grpc_health_v1"
)

type service struct {
	grpchealth.UnimplementedHealthServer
}

var _ grpchealth.HealthServer = (*service)(nil)

func (s *service) Check(context.Context, *grpchealth.HealthCheckRequest) (*grpchealth.HealthCheckResponse, error) {
	return &grpchealth.HealthCheckResponse{
		Status: grpchealth.HealthCheckResponse_SERVING,
	}, nil
}
func (s *service) Watch(*grpchealth.HealthCheckRequest, grpchealth.Health_WatchServer) error {
	return nil
}

// AddSingletonHealthService helper
func AddSingletonHealthService(builder *di.Builder) {
	coreContracts.AddSingletonIHealthServer(builder, reflect.TypeOf(&service{}))
}
