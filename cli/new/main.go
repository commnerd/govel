package new

const MAIN_GO = `
package main

import (
    "Govel"
    "%s/controllers"
)

func main() {
    controllers.Init()
    Govel.Go()
}
`
