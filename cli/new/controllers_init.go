package new

const CONTROLLERS_INIT = `
package controllers

// This file is maintained by Govel.  Please do not manually modify this file.

import "Govel"

func Init() {
	Govel.Register(&RootController{})
}
`
