package new

const ROUTES_WEB = `
package routes

import "Govel/Route"

Route.Get("/", "RootController@index").Name("root")
`
