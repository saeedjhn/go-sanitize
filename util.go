package sanitize

import (
	"reflect"
)

func isPointer(param interface{}) bool {
	return reflect.ValueOf(param).Kind() == reflect.Ptr
}
