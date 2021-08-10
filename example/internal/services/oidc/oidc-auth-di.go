package oidc

import (
	"reflect"

	grpcdotnetgo_di "github.com/fluffy-bunny/grpcdotnetgo/di"
	middleware_auth "github.com/fluffy-bunny/grpcdotnetgo/middleware/auth"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/middleware/oidc"
	services_Logger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	services_oidc "github.com/fluffy-bunny/grpcdotnetgo/services/oidc"
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
