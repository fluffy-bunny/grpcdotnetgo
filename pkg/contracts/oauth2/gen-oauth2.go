// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package oauth2

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeIOAuth2Authenticator used when your service claims to implement IOAuth2Authenticator
var ReflectTypeIOAuth2Authenticator = di.GetInterfaceReflectType((*IOAuth2Authenticator)(nil))

// AddSingletonIOAuth2Authenticator adds a type that implements IOAuth2Authenticator
func AddSingletonIOAuth2Authenticator(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SINGLETON", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIOAuth2AuthenticatorWithMetadata adds a type that implements IOAuth2Authenticator
func AddSingletonIOAuth2AuthenticatorWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SINGLETON", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIOAuth2AuthenticatorByObj adds a prebuilt obj
func AddSingletonIOAuth2AuthenticatorByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SINGLETON", reflect.TypeOf(obj), _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIOAuth2AuthenticatorByObjWithMetadata adds a prebuilt obj
func AddSingletonIOAuth2AuthenticatorByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SINGLETON", reflect.TypeOf(obj), _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIOAuth2AuthenticatorByFunc adds a type by a custom func
func AddSingletonIOAuth2AuthenticatorByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SINGLETON", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIOAuth2AuthenticatorByFuncWithMetadata adds a type by a custom func
func AddSingletonIOAuth2AuthenticatorByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SINGLETON", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIOAuth2Authenticator adds a type that implements IOAuth2Authenticator
func AddTransientIOAuth2Authenticator(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("TRANSIENT", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIOAuth2AuthenticatorWithMetadata adds a type that implements IOAuth2Authenticator
func AddTransientIOAuth2AuthenticatorWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("TRANSIENT", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIOAuth2AuthenticatorByFunc adds a type by a custom func
func AddTransientIOAuth2AuthenticatorByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("TRANSIENT", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIOAuth2AuthenticatorByFuncWithMetadata adds a type by a custom func
func AddTransientIOAuth2AuthenticatorByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("TRANSIENT", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIOAuth2Authenticator adds a type that implements IOAuth2Authenticator
func AddScopedIOAuth2Authenticator(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SCOPED", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIOAuth2AuthenticatorWithMetadata adds a type that implements IOAuth2Authenticator
func AddScopedIOAuth2AuthenticatorWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SCOPED", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIOAuth2AuthenticatorByFunc adds a type by a custom func
func AddScopedIOAuth2AuthenticatorByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SCOPED", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIOAuth2AuthenticatorByFuncWithMetadata adds a type by a custom func
func AddScopedIOAuth2AuthenticatorByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIOAuth2Authenticator)
	_logAddIOAuth2Authenticator("SCOPED", implType, _getImplementedIOAuth2AuthenticatorNames(implementedTypes...),
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIOAuth2AuthenticatorExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIOAuth2Authenticator removes all IOAuth2Authenticator from the DI
func RemoveAllIOAuth2Authenticator(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIOAuth2Authenticator)
}

// GetIOAuth2AuthenticatorFromContainer alternative to SafeGetIOAuth2AuthenticatorFromContainer but panics of object is not present
func GetIOAuth2AuthenticatorFromContainer(ctn di.Container) IOAuth2Authenticator {
	return ctn.GetByType(ReflectTypeIOAuth2Authenticator).(IOAuth2Authenticator)
}

// GetManyIOAuth2AuthenticatorFromContainer alternative to SafeGetManyIOAuth2AuthenticatorFromContainer but panics of object is not present
func GetManyIOAuth2AuthenticatorFromContainer(ctn di.Container) []IOAuth2Authenticator {
	objs := ctn.GetManyByType(ReflectTypeIOAuth2Authenticator)
	var results []IOAuth2Authenticator
	for _, obj := range objs {
		results = append(results, obj.(IOAuth2Authenticator))
	}
	return results
}

// SafeGetIOAuth2AuthenticatorFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIOAuth2AuthenticatorFromContainer(ctn di.Container) (IOAuth2Authenticator, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIOAuth2Authenticator)
	if err != nil {
		return nil, err
	}
	return obj.(IOAuth2Authenticator), nil
}

// GetIOAuth2AuthenticatorDefinition returns that last definition registered that this container can provide
func GetIOAuth2AuthenticatorDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIOAuth2Authenticator)
	return def
}

// GetIOAuth2AuthenticatorDefinitions returns all definitions that this container can provide
func GetIOAuth2AuthenticatorDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIOAuth2Authenticator)
	return defs
}

// SafeGetManyIOAuth2AuthenticatorFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIOAuth2AuthenticatorFromContainer(ctn di.Container) ([]IOAuth2Authenticator, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIOAuth2Authenticator)
	if err != nil {
		return nil, err
	}
	var results []IOAuth2Authenticator
	for _, obj := range objs {
		results = append(results, obj.(IOAuth2Authenticator))
	}
	return results, nil
}

type _logIOAuth2AuthenticatorExtra struct {
	Name  string
	Value interface{}
}

func _logAddIOAuth2Authenticator(scopeType string, implType reflect.Type, interfaces string, extra ..._logIOAuth2AuthenticatorExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIOAuth2AuthenticatorNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
