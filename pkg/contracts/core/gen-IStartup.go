// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package core

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeICoreConfig used when your service claims to implement ICoreConfig
var ReflectTypeICoreConfig = di.GetInterfaceReflectType((*ICoreConfig)(nil))

// AddSingletonICoreConfigByObj adds a prebuilt obj
func AddSingletonICoreConfigByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeICoreConfig)
}

// AddSingletonICoreConfig adds a type that implements ICoreConfig
func AddSingletonICoreConfig(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeICoreConfig)
}

// AddSingletonICoreConfigByFunc adds a type by a custom func
func AddSingletonICoreConfigByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeICoreConfig)
}

// AddTransientICoreConfig adds a type that implements ICoreConfig
func AddTransientICoreConfig(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeICoreConfig)
}

// AddTransientICoreConfigByFunc adds a type by a custom func
func AddTransientICoreConfigByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeICoreConfig)
}

// AddScopedICoreConfig adds a type that implements ICoreConfig
func AddScopedICoreConfig(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeICoreConfig)
}

// AddScopedICoreConfigByFunc adds a type by a custom func
func AddScopedICoreConfigByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeICoreConfig)
}

// RemoveAllICoreConfig removes all ICoreConfig from the DI
func RemoveAllICoreConfig(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeICoreConfig)
}

// GetICoreConfigFromContainer alternative to SafeGetICoreConfigFromContainer but panics of object is not present
func GetICoreConfigFromContainer(ctn di.Container) ICoreConfig {
	return ctn.GetByType(ReflectTypeICoreConfig).(ICoreConfig)
}

// GetManyICoreConfigFromContainer alternative to SafeGetManyICoreConfigFromContainer but panics of object is not present
func GetManyICoreConfigFromContainer(ctn di.Container) []ICoreConfig {
	objs := ctn.GetManyByType(ReflectTypeICoreConfig)
	var results []ICoreConfig
	for _, obj := range objs {
		results = append(results, obj.(ICoreConfig))
	}
	return results
}

// SafeGetICoreConfigFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetICoreConfigFromContainer(ctn di.Container) (ICoreConfig, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeICoreConfig)
	if err != nil {
		return nil, err
	}
	return obj.(ICoreConfig), nil
}

// SafeGetManyICoreConfigFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyICoreConfigFromContainer(ctn di.Container) ([]ICoreConfig, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeICoreConfig)
	if err != nil {
		return nil, err
	}
	var results []ICoreConfig
	for _, obj := range objs {
		results = append(results, obj.(ICoreConfig))
	}
	return results, nil
}

// ReflectTypeIStartup used when your service claims to implement IStartup
var ReflectTypeIStartup = di.GetInterfaceReflectType((*IStartup)(nil))

// AddSingletonIStartupByObj adds a prebuilt obj
func AddSingletonIStartupByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIStartup)
}

// AddSingletonIStartup adds a type that implements IStartup
func AddSingletonIStartup(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIStartup)
}

// AddSingletonIStartupByFunc adds a type by a custom func
func AddSingletonIStartupByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIStartup)
}

// AddTransientIStartup adds a type that implements IStartup
func AddTransientIStartup(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIStartup)
}

// AddTransientIStartupByFunc adds a type by a custom func
func AddTransientIStartupByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIStartup)
}

// AddScopedIStartup adds a type that implements IStartup
func AddScopedIStartup(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIStartup)
}

// AddScopedIStartupByFunc adds a type by a custom func
func AddScopedIStartupByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIStartup)
}

// RemoveAllIStartup removes all IStartup from the DI
func RemoveAllIStartup(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIStartup)
}

// GetIStartupFromContainer alternative to SafeGetIStartupFromContainer but panics of object is not present
func GetIStartupFromContainer(ctn di.Container) IStartup {
	return ctn.GetByType(ReflectTypeIStartup).(IStartup)
}

// GetManyIStartupFromContainer alternative to SafeGetManyIStartupFromContainer but panics of object is not present
func GetManyIStartupFromContainer(ctn di.Container) []IStartup {
	objs := ctn.GetManyByType(ReflectTypeIStartup)
	var results []IStartup
	for _, obj := range objs {
		results = append(results, obj.(IStartup))
	}
	return results
}

// SafeGetIStartupFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIStartupFromContainer(ctn di.Container) (IStartup, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIStartup)
	if err != nil {
		return nil, err
	}
	return obj.(IStartup), nil
}

// SafeGetManyIStartupFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIStartupFromContainer(ctn di.Container) ([]IStartup, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIStartup)
	if err != nil {
		return nil, err
	}
	var results []IStartup
	for _, obj := range objs {
		results = append(results, obj.(IStartup))
	}
	return results, nil
}

// ReflectTypeIUnaryServerInterceptorBuilder used when your service claims to implement IUnaryServerInterceptorBuilder
var ReflectTypeIUnaryServerInterceptorBuilder = di.GetInterfaceReflectType((*IUnaryServerInterceptorBuilder)(nil))

// AddSingletonIUnaryServerInterceptorBuilderByObj adds a prebuilt obj
func AddSingletonIUnaryServerInterceptorBuilderByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIUnaryServerInterceptorBuilder)
}

// AddSingletonIUnaryServerInterceptorBuilder adds a type that implements IUnaryServerInterceptorBuilder
func AddSingletonIUnaryServerInterceptorBuilder(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIUnaryServerInterceptorBuilder)
}

// AddSingletonIUnaryServerInterceptorBuilderByFunc adds a type by a custom func
func AddSingletonIUnaryServerInterceptorBuilderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIUnaryServerInterceptorBuilder)
}

// AddTransientIUnaryServerInterceptorBuilder adds a type that implements IUnaryServerInterceptorBuilder
func AddTransientIUnaryServerInterceptorBuilder(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIUnaryServerInterceptorBuilder)
}

// AddTransientIUnaryServerInterceptorBuilderByFunc adds a type by a custom func
func AddTransientIUnaryServerInterceptorBuilderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIUnaryServerInterceptorBuilder)
}

// AddScopedIUnaryServerInterceptorBuilder adds a type that implements IUnaryServerInterceptorBuilder
func AddScopedIUnaryServerInterceptorBuilder(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIUnaryServerInterceptorBuilder)
}

// AddScopedIUnaryServerInterceptorBuilderByFunc adds a type by a custom func
func AddScopedIUnaryServerInterceptorBuilderByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIUnaryServerInterceptorBuilder)
}

// RemoveAllIUnaryServerInterceptorBuilder removes all IUnaryServerInterceptorBuilder from the DI
func RemoveAllIUnaryServerInterceptorBuilder(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIUnaryServerInterceptorBuilder)
}

// GetIUnaryServerInterceptorBuilderFromContainer alternative to SafeGetIUnaryServerInterceptorBuilderFromContainer but panics of object is not present
func GetIUnaryServerInterceptorBuilderFromContainer(ctn di.Container) IUnaryServerInterceptorBuilder {
	return ctn.GetByType(ReflectTypeIUnaryServerInterceptorBuilder).(IUnaryServerInterceptorBuilder)
}

// GetManyIUnaryServerInterceptorBuilderFromContainer alternative to SafeGetManyIUnaryServerInterceptorBuilderFromContainer but panics of object is not present
func GetManyIUnaryServerInterceptorBuilderFromContainer(ctn di.Container) []IUnaryServerInterceptorBuilder {
	objs := ctn.GetManyByType(ReflectTypeIUnaryServerInterceptorBuilder)
	var results []IUnaryServerInterceptorBuilder
	for _, obj := range objs {
		results = append(results, obj.(IUnaryServerInterceptorBuilder))
	}
	return results
}

// SafeGetIUnaryServerInterceptorBuilderFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIUnaryServerInterceptorBuilderFromContainer(ctn di.Container) (IUnaryServerInterceptorBuilder, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIUnaryServerInterceptorBuilder)
	if err != nil {
		return nil, err
	}
	return obj.(IUnaryServerInterceptorBuilder), nil
}

// SafeGetManyIUnaryServerInterceptorBuilderFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIUnaryServerInterceptorBuilderFromContainer(ctn di.Container) ([]IUnaryServerInterceptorBuilder, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIUnaryServerInterceptorBuilder)
	if err != nil {
		return nil, err
	}
	var results []IUnaryServerInterceptorBuilder
	for _, obj := range objs {
		results = append(results, obj.(IUnaryServerInterceptorBuilder))
	}
	return results, nil
}
