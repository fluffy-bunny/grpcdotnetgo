package utils

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/copier"
)

// PrettyPrintRedacted ...
func PrettyPrintRedacted(source, dest interface{}) {
	MakeRedactedCopy(source, dest)
	fmt.Println(PrettyJSON(dest))
}

// MakeRedactedCopy ...
func MakeRedactedCopy(src, dst interface{}) error {
	err := copier.Copy(dst, src)
	if err != nil {
		return err
	}
	Redact(reflect.ValueOf(dst))
	return nil
}

// Redact ...
func Redact(reflectValue reflect.Value) {
	reflectValue = indirect(reflectValue)

	if reflectValue.Kind() == reflect.Struct {
		l := reflectValue.NumField()
		for i := 0; i < l; i++ {
			field := indirect(reflectValue.Field(i))
			fieldKind := field.Kind()

			if fieldKind == reflect.Struct || fieldKind == reflect.Slice {
				Redact(field)
				continue
			}

			tt := reflectValue.Type().Field(i)
			if hasTag(tt, "redact") {
				if fieldKind == reflect.String {
					field.SetString(redactString(field.String()))
				}
			}
		}
		return
	}

	if reflectValue.Kind() == reflect.Slice {
		l := reflectValue.Len()
		for i := 0; i < l; i++ {
			Redact(reflectValue.Index(i))
		}
		return
	}
}
func redactString(s string) string {
	if len(s) >= 8 {
		return s[:2] + "**REDACTED**" + s[len(s)-2:]
	} else if len(s) >= 6 {
		return s[:2] + "**REDACTED**"
	}
	return "**REDACTED**"
}
func hasTag(val reflect.StructField, tag string) bool {
	c := val.Tag
	g := c.Get(tag)
	return len(g) > 0
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}
func indirectType(v reflect.Type) reflect.Type {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Slice {
		return v.Elem()
	}
	return v
}
