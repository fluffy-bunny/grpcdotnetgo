// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package claimsprincipal

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeIClaimsPrincipal used when your service claims to implement IClaimsPrincipal
var ReflectTypeIClaimsPrincipal = di.GetInterfaceReflectType((*IClaimsPrincipal)(nil))

// AddSingletonIClaimsPrincipalByObj adds a prebuilt obj
func AddSingletonIClaimsPrincipalByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIClaimsPrincipal)
}

// AddSingletonIClaimsPrincipal adds a type that implements IClaimsPrincipal
func AddSingletonIClaimsPrincipal(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIClaimsPrincipal)
}

// AddSingletonIClaimsPrincipalByFunc adds a type by a custom func
func AddSingletonIClaimsPrincipalByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIClaimsPrincipal)
}

// AddTransientIClaimsPrincipal adds a type that implements IClaimsPrincipal
func AddTransientIClaimsPrincipal(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIClaimsPrincipal)
}

// AddTransientIClaimsPrincipalByFunc adds a type by a custom func
func AddTransientIClaimsPrincipalByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIClaimsPrincipal)
}

// AddScopedIClaimsPrincipal adds a type that implements IClaimsPrincipal
func AddScopedIClaimsPrincipal(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIClaimsPrincipal)
}

// AddScopedIClaimsPrincipalByFunc adds a type by a custom func
func AddScopedIClaimsPrincipalByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIClaimsPrincipal)
}

// GetIClaimsPrincipalFromContainer alternative to SafeGetIClaimsPrincipalFromContainer but panics of object is not present
func GetIClaimsPrincipalFromContainer(ctn di.Container) IClaimsPrincipal {
	return ctn.GetByType(ReflectTypeIClaimsPrincipal).(IClaimsPrincipal)
}

// SafeGetIClaimsPrincipalFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIClaimsPrincipalFromContainer(ctn di.Container) (IClaimsPrincipal, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIClaimsPrincipal)
	if err != nil {
		return nil, err
	}
	return obj.(IClaimsPrincipal), nil
}