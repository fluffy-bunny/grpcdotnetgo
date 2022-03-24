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

// AddSingletonIClaimsPrincipal adds a type that implements IClaimsPrincipal
func AddSingletonIClaimsPrincipal(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIClaimsPrincipalWithMetadata adds a type that implements IClaimsPrincipal
func AddSingletonIClaimsPrincipalWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIClaimsPrincipalByObj adds a prebuilt obj
func AddSingletonIClaimsPrincipalByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIClaimsPrincipalByObjWithMetadata adds a prebuilt obj
func AddSingletonIClaimsPrincipalByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIClaimsPrincipalByFunc adds a type by a custom func
func AddSingletonIClaimsPrincipalByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIClaimsPrincipalByFuncWithMetadata adds a type by a custom func
func AddSingletonIClaimsPrincipalByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIClaimsPrincipal adds a type that implements IClaimsPrincipal
func AddTransientIClaimsPrincipal(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIClaimsPrincipalWithMetadata adds a type that implements IClaimsPrincipal
func AddTransientIClaimsPrincipalWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIClaimsPrincipalByFunc adds a type by a custom func
func AddTransientIClaimsPrincipalByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIClaimsPrincipalByFuncWithMetadata adds a type by a custom func
func AddTransientIClaimsPrincipalByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIClaimsPrincipal adds a type that implements IClaimsPrincipal
func AddScopedIClaimsPrincipal(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIClaimsPrincipalWithMetadata adds a type that implements IClaimsPrincipal
func AddScopedIClaimsPrincipalWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIClaimsPrincipalByFunc adds a type by a custom func
func AddScopedIClaimsPrincipalByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIClaimsPrincipalByFuncWithMetadata adds a type by a custom func
func AddScopedIClaimsPrincipalByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIClaimsPrincipal)
	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIClaimsPrincipal removes all IClaimsPrincipal from the DI
func RemoveAllIClaimsPrincipal(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIClaimsPrincipal)
}

// GetIClaimsPrincipalFromContainer alternative to SafeGetIClaimsPrincipalFromContainer but panics of object is not present
func GetIClaimsPrincipalFromContainer(ctn di.Container) IClaimsPrincipal {
	return ctn.GetByType(ReflectTypeIClaimsPrincipal).(IClaimsPrincipal)
}

// GetManyIClaimsPrincipalFromContainer alternative to SafeGetManyIClaimsPrincipalFromContainer but panics of object is not present
func GetManyIClaimsPrincipalFromContainer(ctn di.Container) []IClaimsPrincipal {
	objs := ctn.GetManyByType(ReflectTypeIClaimsPrincipal)
	var results []IClaimsPrincipal
	for _, obj := range objs {
		results = append(results, obj.(IClaimsPrincipal))
	}
	return results
}

// SafeGetIClaimsPrincipalFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIClaimsPrincipalFromContainer(ctn di.Container) (IClaimsPrincipal, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIClaimsPrincipal)
	if err != nil {
		return nil, err
	}
	return obj.(IClaimsPrincipal), nil
}

// GetIClaimsPrincipalDefinition returns that last definition registered that this container can provide
func GetIClaimsPrincipalDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIClaimsPrincipal)
	return def
}

// GetIClaimsPrincipalDefinitions returns all definitions that this container can provide
func GetIClaimsPrincipalDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIClaimsPrincipal)
	return defs
}

// SafeGetManyIClaimsPrincipalFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIClaimsPrincipalFromContainer(ctn di.Container) ([]IClaimsPrincipal, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIClaimsPrincipal)
	if err != nil {
		return nil, err
	}
	var results []IClaimsPrincipal
	for _, obj := range objs {
		results = append(results, obj.(IClaimsPrincipal))
	}
	return results, nil
}
