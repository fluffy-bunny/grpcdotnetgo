// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package oauth2

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeIOauth2 used when your service claims to implement IOauth2
var ReflectTypeIOauth2 = di.GetInterfaceReflectType((*IOauth2)(nil))

// AddSingletonIOauth2ByObj adds a prebuilt obj
func AddSingletonIOauth2ByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIOauth2)
}

// AddSingletonIOauth2 adds a type that implements IOauth2
func AddSingletonIOauth2(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIOauth2)
}

// AddSingletonIOauth2ByFunc adds a type by a custom func
func AddSingletonIOauth2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIOauth2)
}

// AddTransientIOauth2 adds a type that implements IOauth2
func AddTransientIOauth2(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIOauth2)
}

// AddTransientIOauth2ByFunc adds a type by a custom func
func AddTransientIOauth2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIOauth2)
}

// AddScopedIOauth2 adds a type that implements IOauth2
func AddScopedIOauth2(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIOauth2)
}

// AddScopedIOauth2ByFunc adds a type by a custom func
func AddScopedIOauth2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIOauth2)
}

// GetIOauth2FromContainer alternative to SafeGetIOauth2FromContainer but panics of object is not present
func GetIOauth2FromContainer(ctn di.Container) IOauth2 {
	return ctn.GetByType(ReflectTypeIOauth2).(IOauth2)
}

// SafeGetIOauth2FromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIOauth2FromContainer(ctn di.Container) (IOauth2, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIOauth2)
	if err != nil {
		return nil, err
	}
	return obj.(IOauth2), nil
}
