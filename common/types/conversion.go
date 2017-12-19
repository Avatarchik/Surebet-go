package types

import (
	"reflect"
)

func ToInterface(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	length := s.Len()

	ret := make([]interface{}, length)

	for i := 0; i < length; i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func ToType(slice []interface{}, nilEl interface{}) interface{} {
	length := len(slice)

	typ := reflect.TypeOf(nilEl)
	ret := reflect.MakeSlice(reflect.SliceOf(typ), length, length)

	for i := 0; i < length; i++ {
		ret.Index(i).Set(reflect.ValueOf(slice[i]))
	}

	return ret.Interface()
}
