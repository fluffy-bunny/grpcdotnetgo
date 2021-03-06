// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package logger

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeILogger used when your service claims to implement ILogger
var ReflectTypeILogger = di.GetInterfaceReflectType((*ILogger)(nil))

// AddSingletonILogger adds a type that implements ILogger
func AddSingletonILogger(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SINGLETON", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonILoggerWithMetadata adds a type that implements ILogger
func AddSingletonILoggerWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SINGLETON", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logILoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonILoggerByObj adds a prebuilt obj
func AddSingletonILoggerByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SINGLETON", reflect.TypeOf(obj), _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonILoggerByObjWithMetadata adds a prebuilt obj
func AddSingletonILoggerByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SINGLETON", reflect.TypeOf(obj), _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logILoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonILoggerByFunc adds a type by a custom func
func AddSingletonILoggerByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SINGLETON", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonILoggerByFuncWithMetadata adds a type by a custom func
func AddSingletonILoggerByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SINGLETON", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logILoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientILogger adds a type that implements ILogger
func AddTransientILogger(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("TRANSIENT", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientILoggerWithMetadata adds a type that implements ILogger
func AddTransientILoggerWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("TRANSIENT", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logILoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientILoggerByFunc adds a type by a custom func
func AddTransientILoggerByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("TRANSIENT", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientILoggerByFuncWithMetadata adds a type by a custom func
func AddTransientILoggerByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("TRANSIENT", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logILoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedILogger adds a type that implements ILogger
func AddScopedILogger(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SCOPED", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedILoggerWithMetadata adds a type that implements ILogger
func AddScopedILoggerWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SCOPED", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logILoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedILoggerByFunc adds a type by a custom func
func AddScopedILoggerByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SCOPED", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedILoggerByFuncWithMetadata adds a type by a custom func
func AddScopedILoggerByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeILogger)
	_logAddILogger("SCOPED", implType, _getImplementedILoggerNames(implementedTypes...),
		_logILoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logILoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllILogger removes all ILogger from the DI
func RemoveAllILogger(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeILogger)
}

// GetILoggerFromContainer alternative to SafeGetILoggerFromContainer but panics of object is not present
func GetILoggerFromContainer(ctn di.Container) ILogger {
	return ctn.GetByType(ReflectTypeILogger).(ILogger)
}

// GetManyILoggerFromContainer alternative to SafeGetManyILoggerFromContainer but panics of object is not present
func GetManyILoggerFromContainer(ctn di.Container) []ILogger {
	objs := ctn.GetManyByType(ReflectTypeILogger)
	var results []ILogger
	for _, obj := range objs {
		results = append(results, obj.(ILogger))
	}
	return results
}

// SafeGetILoggerFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetILoggerFromContainer(ctn di.Container) (ILogger, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeILogger)
	if err != nil {
		return nil, err
	}
	return obj.(ILogger), nil
}

// GetILoggerDefinition returns that last definition registered that this container can provide
func GetILoggerDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeILogger)
	return def
}

// GetILoggerDefinitions returns all definitions that this container can provide
func GetILoggerDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeILogger)
	return defs
}

// SafeGetManyILoggerFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyILoggerFromContainer(ctn di.Container) ([]ILogger, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeILogger)
	if err != nil {
		return nil, err
	}
	var results []ILogger
	for _, obj := range objs {
		results = append(results, obj.(ILogger))
	}
	return results, nil
}

type _logILoggerExtra struct {
	Name  string
	Value interface{}
}

func _logAddILogger(scopeType string, implType reflect.Type, interfaces string, extra ..._logILoggerExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedILoggerNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

// ReflectTypeISingletonLogger used when your service claims to implement ISingletonLogger
var ReflectTypeISingletonLogger = di.GetInterfaceReflectType((*ISingletonLogger)(nil))

// AddSingletonISingletonLogger adds a type that implements ISingletonLogger
func AddSingletonISingletonLogger(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SINGLETON", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonISingletonLoggerWithMetadata adds a type that implements ISingletonLogger
func AddSingletonISingletonLoggerWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SINGLETON", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logISingletonLoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonISingletonLoggerByObj adds a prebuilt obj
func AddSingletonISingletonLoggerByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SINGLETON", reflect.TypeOf(obj), _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonISingletonLoggerByObjWithMetadata adds a prebuilt obj
func AddSingletonISingletonLoggerByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SINGLETON", reflect.TypeOf(obj), _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logISingletonLoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonISingletonLoggerByFunc adds a type by a custom func
func AddSingletonISingletonLoggerByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SINGLETON", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonISingletonLoggerByFuncWithMetadata adds a type by a custom func
func AddSingletonISingletonLoggerByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SINGLETON", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logISingletonLoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientISingletonLogger adds a type that implements ISingletonLogger
func AddTransientISingletonLogger(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("TRANSIENT", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientISingletonLoggerWithMetadata adds a type that implements ISingletonLogger
func AddTransientISingletonLoggerWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("TRANSIENT", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logISingletonLoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientISingletonLoggerByFunc adds a type by a custom func
func AddTransientISingletonLoggerByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("TRANSIENT", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientISingletonLoggerByFuncWithMetadata adds a type by a custom func
func AddTransientISingletonLoggerByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("TRANSIENT", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logISingletonLoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedISingletonLogger adds a type that implements ISingletonLogger
func AddScopedISingletonLogger(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SCOPED", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedISingletonLoggerWithMetadata adds a type that implements ISingletonLogger
func AddScopedISingletonLoggerWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SCOPED", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logISingletonLoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedISingletonLoggerByFunc adds a type by a custom func
func AddScopedISingletonLoggerByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SCOPED", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedISingletonLoggerByFuncWithMetadata adds a type by a custom func
func AddScopedISingletonLoggerByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeISingletonLogger)
	_logAddISingletonLogger("SCOPED", implType, _getImplementedISingletonLoggerNames(implementedTypes...),
		_logISingletonLoggerExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logISingletonLoggerExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllISingletonLogger removes all ISingletonLogger from the DI
func RemoveAllISingletonLogger(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeISingletonLogger)
}

// GetISingletonLoggerFromContainer alternative to SafeGetISingletonLoggerFromContainer but panics of object is not present
func GetISingletonLoggerFromContainer(ctn di.Container) ISingletonLogger {
	return ctn.GetByType(ReflectTypeISingletonLogger).(ISingletonLogger)
}

// GetManyISingletonLoggerFromContainer alternative to SafeGetManyISingletonLoggerFromContainer but panics of object is not present
func GetManyISingletonLoggerFromContainer(ctn di.Container) []ISingletonLogger {
	objs := ctn.GetManyByType(ReflectTypeISingletonLogger)
	var results []ISingletonLogger
	for _, obj := range objs {
		results = append(results, obj.(ISingletonLogger))
	}
	return results
}

// SafeGetISingletonLoggerFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetISingletonLoggerFromContainer(ctn di.Container) (ISingletonLogger, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeISingletonLogger)
	if err != nil {
		return nil, err
	}
	return obj.(ISingletonLogger), nil
}

// GetISingletonLoggerDefinition returns that last definition registered that this container can provide
func GetISingletonLoggerDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeISingletonLogger)
	return def
}

// GetISingletonLoggerDefinitions returns all definitions that this container can provide
func GetISingletonLoggerDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeISingletonLogger)
	return defs
}

// SafeGetManyISingletonLoggerFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyISingletonLoggerFromContainer(ctn di.Container) ([]ISingletonLogger, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeISingletonLogger)
	if err != nil {
		return nil, err
	}
	var results []ISingletonLogger
	for _, obj := range objs {
		results = append(results, obj.(ISingletonLogger))
	}
	return results, nil
}

type _logISingletonLoggerExtra struct {
	Name  string
	Value interface{}
}

func _logAddISingletonLogger(scopeType string, implType reflect.Type, interfaces string, extra ..._logISingletonLoggerExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedISingletonLoggerNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
