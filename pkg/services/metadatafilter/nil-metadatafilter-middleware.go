package metadatafilter

import (
	"context"
	"reflect"

	contractsmetadatafilter "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/metadatafilter"
	di "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/grpc"
)

type (
	nilMetadataFilterMiddleware struct {
	}
)

// AddSingletonIMetadataFilterMiddlewareNil adds service to the DI container
func AddSingletonIMetadataFilterMiddlewareNil(builder *di.Builder) {
	contractsmetadatafilter.AddSingletonIMetadataFilterMiddleware(builder, reflect.TypeOf(&nilMetadataFilterMiddleware{}))
}

// GetUnaryServerInterceptor ...
func (s *nilMetadataFilterMiddleware) GetUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}
