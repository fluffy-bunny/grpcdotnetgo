package metadatafilter

import (
	"context"
	"reflect"
	"strings"

	contractsmetadatafilter "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/metadatafilter"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
)

type (
	metadataFilterMiddleware struct {
		allowedGlobal       *hashset.StringSet
		allowedByEntryPoint map[string]*hashset.StringSet
	}
)

// AddSingletonIMetadataFilterMiddleware adds service to the DI container
func AddSingletonIMetadataFilterMiddleware(builder *di.Builder,
	allowedGlobal *hashset.StringSet,
	allowedByEntryPoint map[string]*hashset.StringSet) {
	contractsmetadatafilter.AddSingletonIMetadataFilterMiddlewareByFunc(builder, reflect.TypeOf(&metadataFilterMiddleware{}),
		func(ctn di.Container) (interface{}, error) {
			return &metadataFilterMiddleware{
				allowedGlobal:       allowedGlobal,
				allowedByEntryPoint: allowedByEntryPoint,
			}, nil
		})
}

// GetUnaryServerInterceptor ...
func (s *metadataFilterMiddleware) GetUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md := metautils.ExtractIncoming(ctx)

		entryPointAllowed, entryPointExists := s.allowedByEntryPoint[info.FullMethod]
		notAllowedHeaders := []string{}
		for header := range md {
			exists := s.allowedGlobal.Contains(header)
			if exists {
				continue
			}
			// is it explictly allowed for this entry point?
			if entryPointExists {
				exists := entryPointAllowed.Contains(strings.ToLower(header))
				if exists {
					continue
				}
			}
			notAllowedHeaders = append(notAllowedHeaders, header)
		}
		for _, header := range notAllowedHeaders {
			md.Del(header)
		}
		// commit our changes
		newCtx := md.ToIncoming(ctx)
		return handler(newCtx, req)
	}
}
