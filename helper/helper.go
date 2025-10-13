package helper

import "reflect"

func Equal[T any](x, y T) bool {
	return reflect.DeepEqual(x, y)
}
