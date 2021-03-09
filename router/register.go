package router

import (
	"errors"
	"reflect"
	"fmt"
)

func Register(rt route) error {
	rt, ok := instance.Get(rt.GetPath()).(route)
	if !ok {
		return errors.New(fmt.Sprintf("\"%s\" is not a proper route.", ))
	}
	handler, ok := router.Get(rt).(handler)

	if  {
		return errors.New("Route already exists.")
	}
	instance.Set(rt.GetPath(), rt)
	return nil
}