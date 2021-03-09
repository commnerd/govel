package route

import "github.com/commnerd/govel/util"

func (r *route) Except(iMethods ...interface{}) *route {
	sMethods := []string{}

	for _, iMethod := range iMethods {
		switch methods := iMethod.(type) {
		case []string:
			sMethods = methods
		case string:
			sMethods = append(sMethods, methods)
		default:
			panic("Type not allowed.")
		}
	}
	
	for _, method := range sMethods {
		if exists, index := util.InArray(method, r.only); exists {
			r.only = append(r.only[:index], r.only[index+1:]...)
		}
	}

	return r
}