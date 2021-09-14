package di

import (
	"fmt"
	"reflect"

	"github.com/fatih/structtag"
)

// Invoke - firstResult, err := Invoke(AnyStructInterface, MethodName, Params...)
func invoke(any interface{}, name string, args ...interface{}) ([]reflect.Value, error) {
	method := reflect.ValueOf(any).MethodByName(name)
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return nil, fmt.Errorf("method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return nil, fmt.Errorf("method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return nil, fmt.Errorf("method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return nil, fmt.Errorf("method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}

	return method.Call(in), nil
}
func MakeDefaultBuildByType(rtElem reflect.Type, def Def) func(ctn Container) (interface{}, error) {

	objMaker := MakeInjectBuilderFunc(rtElem, def)
	return func(ctn Container) (interface{}, error) {
		rtValue := reflect.New(rtElem)
		dst := rtValue.Interface()
		return objMaker(ctn, dst)
	}
}

// MakeInjectBuilderFunc is EXPENSIVE consider making direct calls to GetByType and GetManyByType directly
func MakeInjectBuilderFunc(rt reflect.Type, def Def) func(ctn Container, dst interface{}) (interface{}, error) {
	setters := []func(ctn Container, dst interface{}){}
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag
		tags, err := structtag.Parse(string(tag))
		if err != nil {
			panic(err)
		}
		if !hasKey("inject", tags) {
			continue
		}
		rtField := field.Type
		fieldName := field.Name
		switch kind := rtField.Kind(); kind {
		case reflect.Array, reflect.Slice:
			sliceElem := rtField.Elem()
			setters = append(setters, func(ctn Container, dst interface{}) {
				v := reflect.ValueOf(dst).Elem()
				f := v.FieldByName(fieldName)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						sliceType := reflect.SliceOf(sliceElem)
						sliceV := reflect.New(sliceType).Elem()
						var objs []interface{}
						if def.SafeInject {
							objs, _ = ctn.SafeGetManyByType(sliceElem)
						} else {
							objs = ctn.GetManyByType(sliceElem)
						}
						if objs != nil {
							for _, obj := range objs {
								tsV := reflect.ValueOf(obj)
								sliceV = reflect.Append(sliceV, tsV)
							}
							f.Set(sliceV)
						}
					}
				}
			})
		default:
			setters = append(setters, func(ctn Container, dst interface{}) {
				v := reflect.ValueOf(dst).Elem()
				f := v.FieldByName(fieldName)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						var obj interface{}
						if def.SafeInject {
							obj, _ = ctn.SafeGetByType(rtField)
						} else {
							obj = ctn.GetByType(rtField)
						}
						if obj != nil {
							objValue := reflect.ValueOf(obj)
							f.Set(objValue)
						}
					}
				}
			})
		}
	}

	return func(ctn Container, dst interface{}) (interface{}, error) {
		rtDst := reflect.TypeOf(dst)
		switch kind := rtDst.Kind(); kind {
		case reflect.Ptr:
			rtElem := rtDst.Elem()
			if rtElem != rt {
				panic("type mismatch")
			}
		default:
			panic("Must be a ptr to a struct. type mismatch")
		}
		for _, setter := range setters {
			setter(ctn, dst)
		}
		if def.hasCtor {
			invoke(dst, "Ctor")
		}
		return dst, nil
	}
}

func hasKey(key string, tags *structtag.Tags) bool {
	for _, k := range tags.Keys() {
		if k == key {
			return true
		}
	}
	return false
}
