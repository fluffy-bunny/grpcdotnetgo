// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package cache

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeICache used when your service claims to implement ICache
var ReflectTypeICache = di.GetInterfaceReflectType((*ICache)(nil))

// AddSingletonICacheByObj adds a prebuilt obj
func AddSingletonICacheByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeICache)
}

// AddSingletonICache adds a type that implements ICache
func AddSingletonICache(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeICache)
}

// AddSingletonICacheByFunc adds a type by a custom func
func AddSingletonICacheByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeICache)
}

// AddTransientICache adds a type that implements ICache
func AddTransientICache(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeICache)
}

// AddTransientICacheByFunc adds a type by a custom func
func AddTransientICacheByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeICache)
}

// AddScopedICache adds a type that implements ICache
func AddScopedICache(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeICache)
}

// AddScopedICacheByFunc adds a type by a custom func
func AddScopedICacheByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeICache)
}

// RemoveAllICache removes all ICache from the DI
func RemoveAllICache(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeICache)
}

// GetICacheFromContainer alternative to SafeGetICacheFromContainer but panics of object is not present
func GetICacheFromContainer(ctn di.Container) ICache {
	return ctn.GetByType(ReflectTypeICache).(ICache)
}

// GetManyICacheFromContainer alternative to SafeGetManyICacheFromContainer but panics of object is not present
func GetManyICacheFromContainer(ctn di.Container) []ICache {
	objs := ctn.GetManyByType(ReflectTypeICache)
	var results []ICache
	for _, obj := range objs {
		results = append(results, obj.(ICache))
	}
	return results
}

// SafeGetICacheFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetICacheFromContainer(ctn di.Container) (ICache, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeICache)
	if err != nil {
		return nil, err
	}
	return obj.(ICache), nil
}

// SafeGetManyICacheFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyICacheFromContainer(ctn di.Container) ([]ICache, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeICache)
	if err != nil {
		return nil, err
	}
	var results []ICache
	for _, obj := range objs {
		results = append(results, obj.(ICache))
	}
	return results, nil
}

// ReflectTypeIMemoryCache used when your service claims to implement IMemoryCache
var ReflectTypeIMemoryCache = di.GetInterfaceReflectType((*IMemoryCache)(nil))

// AddSingletonIMemoryCacheByObj adds a prebuilt obj
func AddSingletonIMemoryCacheByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIMemoryCache)
}

// AddSingletonIMemoryCache adds a type that implements IMemoryCache
func AddSingletonIMemoryCache(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIMemoryCache)
}

// AddSingletonIMemoryCacheByFunc adds a type by a custom func
func AddSingletonIMemoryCacheByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIMemoryCache)
}

// AddTransientIMemoryCache adds a type that implements IMemoryCache
func AddTransientIMemoryCache(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIMemoryCache)
}

// AddTransientIMemoryCacheByFunc adds a type by a custom func
func AddTransientIMemoryCacheByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIMemoryCache)
}

// AddScopedIMemoryCache adds a type that implements IMemoryCache
func AddScopedIMemoryCache(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIMemoryCache)
}

// AddScopedIMemoryCacheByFunc adds a type by a custom func
func AddScopedIMemoryCacheByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIMemoryCache)
}

// RemoveAllIMemoryCache removes all IMemoryCache from the DI
func RemoveAllIMemoryCache(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIMemoryCache)
}

// GetIMemoryCacheFromContainer alternative to SafeGetIMemoryCacheFromContainer but panics of object is not present
func GetIMemoryCacheFromContainer(ctn di.Container) IMemoryCache {
	return ctn.GetByType(ReflectTypeIMemoryCache).(IMemoryCache)
}

// GetManyIMemoryCacheFromContainer alternative to SafeGetManyIMemoryCacheFromContainer but panics of object is not present
func GetManyIMemoryCacheFromContainer(ctn di.Container) []IMemoryCache {
	objs := ctn.GetManyByType(ReflectTypeIMemoryCache)
	var results []IMemoryCache
	for _, obj := range objs {
		results = append(results, obj.(IMemoryCache))
	}
	return results
}

// SafeGetIMemoryCacheFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIMemoryCacheFromContainer(ctn di.Container) (IMemoryCache, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIMemoryCache)
	if err != nil {
		return nil, err
	}
	return obj.(IMemoryCache), nil
}

// SafeGetManyIMemoryCacheFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIMemoryCacheFromContainer(ctn di.Container) ([]IMemoryCache, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIMemoryCache)
	if err != nil {
		return nil, err
	}
	var results []IMemoryCache
	for _, obj := range objs {
		results = append(results, obj.(IMemoryCache))
	}
	return results, nil
}