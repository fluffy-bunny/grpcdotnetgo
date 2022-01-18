// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package serviceprovider

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeIServiceProvider used when your service claims to implement IServiceProvider
var ReflectTypeIServiceProvider = di.GetInterfaceReflectType((*IServiceProvider)(nil))

// AddSingletonIServiceProviderByObj adds a prebuilt obj
func AddSingletonIServiceProviderByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIServiceProvider)
}

// AddSingletonIServiceProvider adds a type that implements IServiceProvider
func AddSingletonIServiceProvider(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIServiceProvider)
}

// AddSingletonIServiceProviderByFunc adds a type by a custom func
func AddSingletonIServiceProviderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIServiceProvider)
}

// AddTransientIServiceProvider adds a type that implements IServiceProvider
func AddTransientIServiceProvider(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIServiceProvider)
}

// AddTransientIServiceProviderByFunc adds a type by a custom func
func AddTransientIServiceProviderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIServiceProvider)
}

// AddScopedIServiceProvider adds a type that implements IServiceProvider
func AddScopedIServiceProvider(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIServiceProvider)
}

// AddScopedIServiceProviderByFunc adds a type by a custom func
func AddScopedIServiceProviderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIServiceProvider)
}

// RemoveAllIServiceProvider removes all IServiceProvider from the DI
func RemoveAllIServiceProvider(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIServiceProvider)
}

// GetIServiceProviderFromContainer alternative to SafeGetIServiceProviderFromContainer but panics of object is not present
func GetIServiceProviderFromContainer(ctn di.Container) IServiceProvider {
	return ctn.GetByType(ReflectTypeIServiceProvider).(IServiceProvider)
}

// GetManyIServiceProviderFromContainer alternative to SafeGetManyIServiceProviderFromContainer but panics of object is not present
func GetManyIServiceProviderFromContainer(ctn di.Container) []IServiceProvider {
	objs := ctn.GetManyByType(ReflectTypeIServiceProvider)
	var results []IServiceProvider
	for _, obj := range objs {
		results = append(results, obj.(IServiceProvider))
	}
	return results
}

// SafeGetIServiceProviderFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIServiceProviderFromContainer(ctn di.Container) (IServiceProvider, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIServiceProvider)
	if err != nil {
		return nil, err
	}
	return obj.(IServiceProvider), nil
}

// SafeGetManyIServiceProviderFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIServiceProviderFromContainer(ctn di.Container) ([]IServiceProvider, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIServiceProvider)
	if err != nil {
		return nil, err
	}
	var results []IServiceProvider
	for _, obj := range objs {
		results = append(results, obj.(IServiceProvider))
	}
	return results, nil
}

// ReflectTypeISingletonServiceProvider used when your service claims to implement ISingletonServiceProvider
var ReflectTypeISingletonServiceProvider = di.GetInterfaceReflectType((*ISingletonServiceProvider)(nil))

// AddSingletonISingletonServiceProviderByObj adds a prebuilt obj
func AddSingletonISingletonServiceProviderByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeISingletonServiceProvider)
}

// AddSingletonISingletonServiceProvider adds a type that implements ISingletonServiceProvider
func AddSingletonISingletonServiceProvider(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeISingletonServiceProvider)
}

// AddSingletonISingletonServiceProviderByFunc adds a type by a custom func
func AddSingletonISingletonServiceProviderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISingletonServiceProvider)
}

// AddTransientISingletonServiceProvider adds a type that implements ISingletonServiceProvider
func AddTransientISingletonServiceProvider(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeISingletonServiceProvider)
}

// AddTransientISingletonServiceProviderByFunc adds a type by a custom func
func AddTransientISingletonServiceProviderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISingletonServiceProvider)
}

// AddScopedISingletonServiceProvider adds a type that implements ISingletonServiceProvider
func AddScopedISingletonServiceProvider(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeISingletonServiceProvider)
}

// AddScopedISingletonServiceProviderByFunc adds a type by a custom func
func AddScopedISingletonServiceProviderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISingletonServiceProvider)
}

// RemoveAllISingletonServiceProvider removes all ISingletonServiceProvider from the DI
func RemoveAllISingletonServiceProvider(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeISingletonServiceProvider)
}

// GetISingletonServiceProviderFromContainer alternative to SafeGetISingletonServiceProviderFromContainer but panics of object is not present
func GetISingletonServiceProviderFromContainer(ctn di.Container) ISingletonServiceProvider {
	return ctn.GetByType(ReflectTypeISingletonServiceProvider).(ISingletonServiceProvider)
}

// GetManyISingletonServiceProviderFromContainer alternative to SafeGetManyISingletonServiceProviderFromContainer but panics of object is not present
func GetManyISingletonServiceProviderFromContainer(ctn di.Container) []ISingletonServiceProvider {
	objs := ctn.GetManyByType(ReflectTypeISingletonServiceProvider)
	var results []ISingletonServiceProvider
	for _, obj := range objs {
		results = append(results, obj.(ISingletonServiceProvider))
	}
	return results
}

// SafeGetISingletonServiceProviderFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetISingletonServiceProviderFromContainer(ctn di.Container) (ISingletonServiceProvider, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeISingletonServiceProvider)
	if err != nil {
		return nil, err
	}
	return obj.(ISingletonServiceProvider), nil
}

// SafeGetManyISingletonServiceProviderFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyISingletonServiceProviderFromContainer(ctn di.Container) ([]ISingletonServiceProvider, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeISingletonServiceProvider)
	if err != nil {
		return nil, err
	}
	var results []ISingletonServiceProvider
	for _, obj := range objs {
		results = append(results, obj.(ISingletonServiceProvider))
	}
	return results, nil
}
