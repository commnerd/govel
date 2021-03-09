package new

const MAIN_GO = `package main

import (
    "govel"
    "%s/controllers"
)

func main() {
    controllers.Init()
    govel.Go()
}
`
