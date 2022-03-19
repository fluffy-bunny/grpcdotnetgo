package logger

import (
	"context"
	"strings"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/wellknown"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// EnsureCorrelationIDUnaryServerInterceptor ...
func EnsureCorrelationIDUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var correlationID string // if not found in header, we generate a new one
		var requestID = "0"
		md := metautils.ExtractIncoming(ctx)
		var loggerMap = make(map[string]string)

		for key, v := range md {
			lowerKey := strings.ToLower(key)
			if lowerKey == wellknown.XCorrelationIDName {
				correlationID = v[0]
			}
			if lowerKey == wellknown.XRequestID {
				requestID = v[0]
			}
		}

		if len(correlationID) == 0 {
			correlationID = genCorrelationID()
			md[wellknown.XCorrelationIDName] = []string{correlationID}
		}

		loggerMap["correlation_id"] = correlationID
		// this came into us, so its a parent
		items := md[wellknown.XSpanName]
		if items != nil && len(items) > 0 {
			loggerMap[wellknown.LogParentName] = items[0]
			md[wellknown.XParentName] = []string{items[0]}
		}
		// generate a new span for this context
		newSpanID := generateSpanID()
		md[wellknown.XSpanName] = []string{newSpanID}
		loggerMap[wellknown.LogSpanName] = newSpanID
		log := zerolog.Ctx(ctx)
		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			for k, v := range loggerMap {
				c = c.Str(k, v)
			}
			return c
		})
		// Return the cleansed metadata context
		ctx = md.ToIncoming(ctx)

		md2 := metadata.Pairs(
			wellknown.XRequestID, requestID,
			wellknown.XCorrelationIDName, correlationID)
		grpc.SendHeader(ctx, md2)
		return handler(ctx, req)
	}
}
func generateSpanID() string {
	return xid.New().String()
}
func genCorrelationID() string {
	return xid.New().String()
}
