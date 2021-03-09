package gerror

import (
	"reflect"
	"errors"
	"fmt"
)

func Throw(iError interface{}) {
	switch err := iError.(type) {
	case error:
		panic(err)
	case string:
		panic(errors.New(err))
	case gError:
		Throw(err.Error())
	default:
		template := "Type \"%s\" cannot be used as an error. Please implement an \"Error() error\" method on this struct to squelch this error."
		panic(errors.New(fmt.Sprintf(template, reflect.TypeOf(iError).Name())))
	}
}