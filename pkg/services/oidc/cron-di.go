package oidc

import (
	"reflect"

	backgroundtasksContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
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
			obj := &serviceJobProvider{
				OIDCConfigAccessor: middleware_oidc.GetOIDCConfigAccessorFromContainer(ctn),
				Storage:            GetOidcBackgroundStorageFromContainer(ctn),
				Logger:             contracts_logger.GetILoggerFromContainer(ctn),
			}
			return obj, nil
		})

	addOidcBackgroundStorage(builder)
}

// GetOidcBackgroundStorageFromContainer helper
func GetOidcBackgroundStorageFromContainer(ctn di.Container) IOidcBackgroundStorage {
	obj := ctn.GetByType(TypeIOidcBackgroundStorage).(IOidcBackgroundStorage)
	return obj
}

// addOidcBackgroundStorage adds service to the DI container
func addOidcBackgroundStorage(builder *di.Builder) {
	log.Info().
		Msg("IoC: addOidcBackgroundStorage")
	types := di.NewTypeSet()
	types.Add(TypeIOidcBackgroundStorage)

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&oidcBackgroundStorage{}),
		Build: func(ctn di.Container) (interface{}, error) {
			obj := &oidcBackgroundStorage{}
			return obj, nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	})
}
