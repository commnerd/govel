package new

const ROUTES_WEB = `package routes

import "govel/route"

func init() {
	route.Get("/", "RootController@index").Name("root")
}
`
