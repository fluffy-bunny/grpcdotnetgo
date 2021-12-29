package oidc

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oidc"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddCronOidcJobProvider adds service to the DI container
func AddCronOidcJobProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddCronOidcJobProvider")

	backgroundtasksContracts.AddSingletonIJobsProviderByFunc(builder,
		reflect.TypeOf(&service{}), func(ctn di.Container) (interface{}, error) {
			obj := &service{
				OIDCConfigAccessor: middleware_oidc.GetOIDCConfigAccessorFromContainer(ctn),
				Storage:            contracts_oidc.GetIOidcBackgroundStorageFromContainer(ctn),
				Logger:             contracts_logger.GetISingletonLoggerFromContainer(ctn),
			}
			return obj, nil
		})

	addOidcBackgroundStorage(builder)
}

// addOidcBackgroundStorage adds service to the DI container
func addOidcBackgroundStorage(builder *di.Builder) {
	log.Info().
		Msg("IoC: addOidcBackgroundStorage")
	contracts_oidc.AddSingletonIOidcBackgroundStorage(builder, reflect.TypeOf(&oidcBackgroundStorage{}))
}
