package app

import "reflect"

func (a *application) Register(i interface{}) error {
	a.Set(reflect.TypeOf(i).String(), i)
	return nil
}
