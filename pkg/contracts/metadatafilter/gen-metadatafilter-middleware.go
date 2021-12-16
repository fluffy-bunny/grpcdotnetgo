// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package metadatafilter

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeIMetadataFilterMiddleware used when your service claims to implement IMetadataFilterMiddleware
var ReflectTypeIMetadataFilterMiddleware = di.GetInterfaceReflectType((*IMetadataFilterMiddleware)(nil))

// AddSingletonIMetadataFilterMiddlewareByObj adds a prebuilt obj
func AddSingletonIMetadataFilterMiddlewareByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIMetadataFilterMiddleware)
}

// AddSingletonIMetadataFilterMiddleware adds a type that implements IMetadataFilterMiddleware
func AddSingletonIMetadataFilterMiddleware(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIMetadataFilterMiddleware)
}

// AddSingletonIMetadataFilterMiddlewareByFunc adds a type by a custom func
func AddSingletonIMetadataFilterMiddlewareByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIMetadataFilterMiddleware)
}

// AddTransientIMetadataFilterMiddleware adds a type that implements IMetadataFilterMiddleware
func AddTransientIMetadataFilterMiddleware(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIMetadataFilterMiddleware)
}

// AddTransientIMetadataFilterMiddlewareByFunc adds a type by a custom func
func AddTransientIMetadataFilterMiddlewareByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIMetadataFilterMiddleware)
}

// AddScopedIMetadataFilterMiddleware adds a type that implements IMetadataFilterMiddleware
func AddScopedIMetadataFilterMiddleware(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIMetadataFilterMiddleware)
}

// AddScopedIMetadataFilterMiddlewareByFunc adds a type by a custom func
func AddScopedIMetadataFilterMiddlewareByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIMetadataFilterMiddleware)
}

// RemoveAllIMetadataFilterMiddleware removes all IMetadataFilterMiddleware from the DI
func RemoveAllIMetadataFilterMiddleware(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIMetadataFilterMiddleware)
}

// GetIMetadataFilterMiddlewareFromContainer alternative to SafeGetIMetadataFilterMiddlewareFromContainer but panics of object is not present
func GetIMetadataFilterMiddlewareFromContainer(ctn di.Container) IMetadataFilterMiddleware {
	return ctn.GetByType(ReflectTypeIMetadataFilterMiddleware).(IMetadataFilterMiddleware)
}

// SafeGetIMetadataFilterMiddlewareFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIMetadataFilterMiddlewareFromContainer(ctn di.Container) (IMetadataFilterMiddleware, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIMetadataFilterMiddleware)
	if err != nil {
		return nil, err
	}
	return obj.(IMetadataFilterMiddleware), nil
}
