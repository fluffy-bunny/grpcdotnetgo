package oidc

import (
	"reflect"

	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/middleware/oidc"
	servicesBackgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/services/backgroundtasks"
	servicesLogger "github.com/fluffy-bunny/grpcdotnetgo/services/logger"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddCronOidcJobProvider adds service to the DI container
func AddCronOidcJobProvider(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddCronOidcJobProvider")
	types := di.NewTypeSet()
	types.Add(servicesBackgroundtasks.TypeIJobsProvider)

	builder.Add(di.Def{
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&service{}),
		Build: func(ctn di.Container) (interface{}, error) {
			obj := &service{
				OIDCConfigAccessor: middleware_oidc.GetOIDCConfigAccessorFromContainer(ctn),
				Storage:            GetOidcBackgroundStorageFromContainer(ctn),
				Logger:             servicesLogger.GetSingletonLoggerFromContainer(ctn),
			}
			return obj, nil
		},
		Close: func(obj interface{}) error {

			return nil
		},
	})
	addOidcBackgroundStorage(builder)

}

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
