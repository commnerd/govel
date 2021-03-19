package new

const ROUTES_WEB = `package routes

import "govel/route"

func web() {
	route.Get("/", "RootController@index").Name("root")
}
`
