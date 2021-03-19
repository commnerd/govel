package new

const BOOTSTRAP_GO = `package main

import (
    "govel"
    "govel/config"
    "%s/controllers"
    "%s/routes"
)

func main() {
    govel.Go()
}
`
