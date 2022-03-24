// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package request

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeIRequest used when your service claims to implement IRequest
var ReflectTypeIRequest = di.GetInterfaceReflectType((*IRequest)(nil))

// AddSingletonIRequest adds a type that implements IRequest
func AddSingletonIRequest(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIRequestWithMetadata adds a type that implements IRequest
func AddSingletonIRequestWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIRequestByObj adds a prebuilt obj
func AddSingletonIRequestByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIRequestByObjWithMetadata adds a prebuilt obj
func AddSingletonIRequestByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIRequestByFunc adds a type by a custom func
func AddSingletonIRequestByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIRequestByFuncWithMetadata adds a type by a custom func
func AddSingletonIRequestByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIRequest adds a type that implements IRequest
func AddTransientIRequest(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIRequestWithMetadata adds a type that implements IRequest
func AddTransientIRequestWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIRequestByFunc adds a type by a custom func
func AddTransientIRequestByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIRequestByFuncWithMetadata adds a type by a custom func
func AddTransientIRequestByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIRequest adds a type that implements IRequest
func AddScopedIRequest(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIRequestWithMetadata adds a type that implements IRequest
func AddScopedIRequestWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIRequestByFunc adds a type by a custom func
func AddScopedIRequestByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIRequestByFuncWithMetadata adds a type by a custom func
func AddScopedIRequestByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIRequest)
	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIRequest removes all IRequest from the DI
func RemoveAllIRequest(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIRequest)
}

// GetIRequestFromContainer alternative to SafeGetIRequestFromContainer but panics of object is not present
func GetIRequestFromContainer(ctn di.Container) IRequest {
	return ctn.GetByType(ReflectTypeIRequest).(IRequest)
}

// GetManyIRequestFromContainer alternative to SafeGetManyIRequestFromContainer but panics of object is not present
func GetManyIRequestFromContainer(ctn di.Container) []IRequest {
	objs := ctn.GetManyByType(ReflectTypeIRequest)
	var results []IRequest
	for _, obj := range objs {
		results = append(results, obj.(IRequest))
	}
	return results
}

// SafeGetIRequestFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIRequestFromContainer(ctn di.Container) (IRequest, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIRequest)
	if err != nil {
		return nil, err
	}
	return obj.(IRequest), nil
}

// GetIRequestDefinition returns that last definition registered that this container can provide
func GetIRequestDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIRequest)
	return def
}

// GetIRequestDefinitions returns all definitions that this container can provide
func GetIRequestDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIRequest)
	return defs
}

// SafeGetManyIRequestFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIRequestFromContainer(ctn di.Container) ([]IRequest, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIRequest)
	if err != nil {
		return nil, err
	}
	var results []IRequest
	for _, obj := range objs {
		results = append(results, obj.(IRequest))
	}
	return results, nil
}

// ReflectTypeIItems used when your service claims to implement IItems
var ReflectTypeIItems = di.GetInterfaceReflectType((*IItems)(nil))

// AddSingletonIItems adds a type that implements IItems
func AddSingletonIItems(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIItemsWithMetadata adds a type that implements IItems
func AddSingletonIItemsWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIItemsByObj adds a prebuilt obj
func AddSingletonIItemsByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIItemsByObjWithMetadata adds a prebuilt obj
func AddSingletonIItemsByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIItemsByFunc adds a type by a custom func
func AddSingletonIItemsByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIItemsByFuncWithMetadata adds a type by a custom func
func AddSingletonIItemsByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIItems adds a type that implements IItems
func AddTransientIItems(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIItemsWithMetadata adds a type that implements IItems
func AddTransientIItemsWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIItemsByFunc adds a type by a custom func
func AddTransientIItemsByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIItemsByFuncWithMetadata adds a type by a custom func
func AddTransientIItemsByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIItems adds a type that implements IItems
func AddScopedIItems(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIItemsWithMetadata adds a type that implements IItems
func AddScopedIItemsWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIItemsByFunc adds a type by a custom func
func AddScopedIItemsByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIItemsByFuncWithMetadata adds a type by a custom func
func AddScopedIItemsByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIItems)
	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIItems removes all IItems from the DI
func RemoveAllIItems(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIItems)
}

// GetIItemsFromContainer alternative to SafeGetIItemsFromContainer but panics of object is not present
func GetIItemsFromContainer(ctn di.Container) IItems {
	return ctn.GetByType(ReflectTypeIItems).(IItems)
}

// GetManyIItemsFromContainer alternative to SafeGetManyIItemsFromContainer but panics of object is not present
func GetManyIItemsFromContainer(ctn di.Container) []IItems {
	objs := ctn.GetManyByType(ReflectTypeIItems)
	var results []IItems
	for _, obj := range objs {
		results = append(results, obj.(IItems))
	}
	return results
}

// SafeGetIItemsFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIItemsFromContainer(ctn di.Container) (IItems, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIItems)
	if err != nil {
		return nil, err
	}
	return obj.(IItems), nil
}

// GetIItemsDefinition returns that last definition registered that this container can provide
func GetIItemsDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIItems)
	return def
}

// GetIItemsDefinitions returns all definitions that this container can provide
func GetIItemsDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIItems)
	return defs
}

// SafeGetManyIItemsFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIItemsFromContainer(ctn di.Container) ([]IItems, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIItems)
	if err != nil {
		return nil, err
	}
	var results []IItems
	for _, obj := range objs {
		results = append(results, obj.(IItems))
	}
	return results, nil
}
