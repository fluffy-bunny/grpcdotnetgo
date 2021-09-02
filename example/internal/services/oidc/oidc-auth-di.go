package oidc

import (
	"reflect"

	grpcdotnetgo_di "github.com/fluffy-bunny/grpcdotnetgo/pkg/di"
	middleware_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/auth"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	services_Logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/logger"
	services_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/oidc"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// AddOIDCAuthHandler adds service to the DI container
func AddOIDCAuthHandler(builder *di.Builder) {
	grpcdotnetgo_di.AddSingletonByType(
		builder,
		func(ctn di.Container) (interface{}, error) {
			accessor := middleware_oidc.GetOIDCConfigAccessorFromContainer(ctn)
			return &service{
				Storage:    services_oidc.GetOidcBackgroundStorageFromContainer(ctn),
				Logger:     services_Logger.GetSingletonLoggerFromContainer(ctn),
				OIDCConfig: accessor.GetOIDCConfig(),
			}, nil
		},
		nil,
		reflect.TypeOf(&service{}),
		middleware_auth.TypeIAuthFuncAccessor)

}
