package reflect

import "reflect"

func GetInterfaceReflectType(i interface{}) reflect.Type {
	return reflect.TypeOf(i).Elem()
}
