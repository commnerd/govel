package new

const ROUTES_WEB = `
package routes

import "Govel/route"

func init() {
	route.Get("/", "RootController@index").Name("root")
}
`
