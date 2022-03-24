// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package timeutils

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeITimeUtils used when your service claims to implement ITimeUtils
var ReflectTypeITimeUtils = di.GetInterfaceReflectType((*ITimeUtils)(nil))

// AddSingletonITimeUtils adds a type that implements ITimeUtils
func AddSingletonITimeUtils(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonITimeUtilsWithMetadata adds a type that implements ITimeUtils
func AddSingletonITimeUtilsWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonITimeUtilsByObj adds a prebuilt obj
func AddSingletonITimeUtilsByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonITimeUtilsByObjWithMetadata adds a prebuilt obj
func AddSingletonITimeUtilsByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonITimeUtilsByFunc adds a type by a custom func
func AddSingletonITimeUtilsByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonITimeUtilsByFuncWithMetadata adds a type by a custom func
func AddSingletonITimeUtilsByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientITimeUtils adds a type that implements ITimeUtils
func AddTransientITimeUtils(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientITimeUtilsWithMetadata adds a type that implements ITimeUtils
func AddTransientITimeUtilsWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientITimeUtilsByFunc adds a type by a custom func
func AddTransientITimeUtilsByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientITimeUtilsByFuncWithMetadata adds a type by a custom func
func AddTransientITimeUtilsByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedITimeUtils adds a type that implements ITimeUtils
func AddScopedITimeUtils(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedITimeUtilsWithMetadata adds a type that implements ITimeUtils
func AddScopedITimeUtilsWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedITimeUtilsByFunc adds a type by a custom func
func AddScopedITimeUtilsByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedITimeUtilsByFuncWithMetadata adds a type by a custom func
func AddScopedITimeUtilsByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITimeUtils)
	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllITimeUtils removes all ITimeUtils from the DI
func RemoveAllITimeUtils(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeITimeUtils)
}

// GetITimeUtilsFromContainer alternative to SafeGetITimeUtilsFromContainer but panics of object is not present
func GetITimeUtilsFromContainer(ctn di.Container) ITimeUtils {
	return ctn.GetByType(ReflectTypeITimeUtils).(ITimeUtils)
}

// GetManyITimeUtilsFromContainer alternative to SafeGetManyITimeUtilsFromContainer but panics of object is not present
func GetManyITimeUtilsFromContainer(ctn di.Container) []ITimeUtils {
	objs := ctn.GetManyByType(ReflectTypeITimeUtils)
	var results []ITimeUtils
	for _, obj := range objs {
		results = append(results, obj.(ITimeUtils))
	}
	return results
}

// SafeGetITimeUtilsFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetITimeUtilsFromContainer(ctn di.Container) (ITimeUtils, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeITimeUtils)
	if err != nil {
		return nil, err
	}
	return obj.(ITimeUtils), nil
}

// GetITimeUtilsDefinition returns that last definition registered that this container can provide
func GetITimeUtilsDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeITimeUtils)
	return def
}

// GetITimeUtilsDefinitions returns all definitions that this container can provide
func GetITimeUtilsDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeITimeUtils)
	return defs
}

// SafeGetManyITimeUtilsFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyITimeUtilsFromContainer(ctn di.Container) ([]ITimeUtils, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeITimeUtils)
	if err != nil {
		return nil, err
	}
	var results []ITimeUtils
	for _, obj := range objs {
		results = append(results, obj.(ITimeUtils))
	}
	return results, nil
}

// ReflectTypeITime used when your service claims to implement ITime
var ReflectTypeITime = di.GetInterfaceReflectType((*ITime)(nil))

// AddSingletonITime adds a type that implements ITime
func AddSingletonITime(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonITimeWithMetadata adds a type that implements ITime
func AddSingletonITimeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonITimeByObj adds a prebuilt obj
func AddSingletonITimeByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonITimeByObjWithMetadata adds a prebuilt obj
func AddSingletonITimeByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonITimeByFunc adds a type by a custom func
func AddSingletonITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonITimeByFuncWithMetadata adds a type by a custom func
func AddSingletonITimeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientITime adds a type that implements ITime
func AddTransientITime(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientITimeWithMetadata adds a type that implements ITime
func AddTransientITimeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientITimeByFunc adds a type by a custom func
func AddTransientITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientITimeByFuncWithMetadata adds a type by a custom func
func AddTransientITimeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedITime adds a type that implements ITime
func AddScopedITime(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedITimeWithMetadata adds a type that implements ITime
func AddScopedITimeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedITimeByFunc adds a type by a custom func
func AddScopedITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedITimeByFuncWithMetadata adds a type by a custom func
func AddScopedITimeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllITime removes all ITime from the DI
func RemoveAllITime(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeITime)
}

// GetITimeFromContainer alternative to SafeGetITimeFromContainer but panics of object is not present
func GetITimeFromContainer(ctn di.Container) ITime {
	return ctn.GetByType(ReflectTypeITime).(ITime)
}

// GetManyITimeFromContainer alternative to SafeGetManyITimeFromContainer but panics of object is not present
func GetManyITimeFromContainer(ctn di.Container) []ITime {
	objs := ctn.GetManyByType(ReflectTypeITime)
	var results []ITime
	for _, obj := range objs {
		results = append(results, obj.(ITime))
	}
	return results
}

// SafeGetITimeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetITimeFromContainer(ctn di.Container) (ITime, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeITime)
	if err != nil {
		return nil, err
	}
	return obj.(ITime), nil
}

// GetITimeDefinition returns that last definition registered that this container can provide
func GetITimeDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeITime)
	return def
}

// GetITimeDefinitions returns all definitions that this container can provide
func GetITimeDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeITime)
	return defs
}

// SafeGetManyITimeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyITimeFromContainer(ctn di.Container) ([]ITime, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeITime)
	if err != nil {
		return nil, err
	}
	var results []ITime
	for _, obj := range objs {
		results = append(results, obj.(ITime))
	}
	return results, nil
}
