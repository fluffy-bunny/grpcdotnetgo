package oidc

import (
	"reflect"

	grpcdotnetgodi "github.com/fluffy-bunny/grpcdotnetgo/di"
	di "github.com/fluffy-bunny/sarulabsdi"
)

func GetOIDCConfigAccessorFromContainer(ctn di.Container) IOIDCConfigAccessor {
	obj := ctn.GetByType(TypeIOIDCConfigAccessor).(IOIDCConfigAccessor)
	return obj
}

// AddOIDCConfigAccessor adds service to the DI container
func AddOIDCConfigAccessor(builder *di.Builder, obj interface{}) {
	grpcdotnetgodi.AddByType(
		builder,
		di.App,
		func(ctn di.Container) (interface{}, error) {
			return obj, nil
		},
		nil,
		reflect.TypeOf(obj),
		TypeIOIDCConfigAccessor)

}
