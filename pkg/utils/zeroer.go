package utils

import "reflect"

// Zeroer :
type Zeroer interface {
	IsZero() bool
}

/*
IsZero : determine is zero value
1. string => empty string
2. bool => false
3. function => nil
4. map => nil (uninitialized map)
5. slice => nil (uninitialized slice)
*/
func IsZero(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	it := v.Interface()
	x, ok := it.(Zeroer)
	if ok {
		return x.IsZero()
	}

	switch v.Kind() {
	case reflect.Interface, reflect.Func:
		return v.IsNil()
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	case reflect.Array:
		if v.Len() == 0 {
			return true
		}
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && IsZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && IsZero(v.Field(i))
		}
		return z
	}

	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return it == z.Interface()
}

// IsEmptyOrNil checks if a value is empty or nil, useful for strings and arrays
func IsEmptyOrNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		if reflect.ValueOf(i).IsNil() {
			return true
		}
		return IsZero(reflect.ValueOf(i))
	case reflect.String:
		return i == ""
	}
	return false //everything else here is a primitive
}

// IsNil is a wholistic nil checker
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false //everything else here is a primitive
}
