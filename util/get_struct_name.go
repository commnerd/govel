package util

import "reflect"

func GetStructName(i interface{}) string {
	return reflect.TypeOf(i).Name()
}