package metadatafilter

import (
	"context"
	"reflect"

	contractsmetadatafilter "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/metadatafilter"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type (
	metadataFilterMiddleware struct {
		alwaysAllowed          map[string]bool
		additionalByEntryPoint map[string]map[string]bool
	}
)

// AddSingletonIMetadataFilterMiddleware adds service to the DI container
func AddSingletonIMetadataFilterMiddleware(builder *di.Builder,
	alwaysAllowed map[string]bool,
	additionalByEntryPoint map[string]map[string]bool) {
	log.Info().
		Msg("IoC: AddSingletonIMetadataFilterMiddleware")
	contractsmetadatafilter.AddSingletonIMetadataFilterMiddlewareByFunc(builder, reflect.TypeOf(&metadataFilterMiddleware{}),
		func(ctn di.Container) (interface{}, error) {
			return &metadataFilterMiddleware{
				alwaysAllowed:          alwaysAllowed,
				additionalByEntryPoint: additionalByEntryPoint,
			}, nil
		})
}

// GetUnaryServerInterceptor ...
func (s *metadataFilterMiddleware) GetUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md := metautils.ExtractIncoming(ctx)

		entryPointAllowed, entryPointExists := s.additionalByEntryPoint[info.FullMethod]
		notAllowedHeaders := []string{}
		for header := range md {
			_, exists := s.alwaysAllowed[header]
			if exists {
				continue
			}
			// is it explictly allowed for this entry point?
			if entryPointExists && entryPointAllowed != nil {
				_, exists := entryPointAllowed[header]
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
