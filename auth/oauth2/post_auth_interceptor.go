package oauth2

import (
	"context"

	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/middleware/dicontext"
	services_logger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	"github.com/gogo/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func FinalAuthVerificationMiddleware(oauth2Context *OAuth2Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestContainer := middleware_dicontext.GetRequestContainer(ctx)
		logger := services_logger.GetScopedLoggerFromContainer(requestContainer)
		loggerZ := logger.GetLogger()
		subLogger := loggerZ.With().Str("FullMethod", info.FullMethod).Logger()

		permissionDeniedFunc := func() (interface{}, error) {
			logger.DebugL(&subLogger).Msg("")
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
		data := ctx.Value(CtxClaimsPrincipalKey)
		if data == nil {
			return permissionDeniedFunc()
		}

		claimsPrincipal := data.(*ClaimsPrincipal)
		elem, ok := oauth2Context.Config.FullMethodNameToClaims[info.FullMethod]
		if ok {
			for _, v := range elem.AND {
				p, ok := claimsPrincipal.FastMap[v.Type]
				if !ok {
					return permissionDeniedFunc()
				}
				_, ok = p[v.Value]
				if !ok {
					return permissionDeniedFunc()
				}
			}

			if elem.OR != nil && len(elem.OR) > 0 {
				var found bool = false
				for _, v := range elem.OR {
					p, ok := claimsPrincipal.FastMap[v.Type]
					if !ok {
						continue
					}
					_, ok = p[v.Value]
					if !ok {
						continue
					}
					found = true
					break
				}
				if !found {
					return permissionDeniedFunc()
				}
			}

		}

		return handler(ctx, req)
	}

}
