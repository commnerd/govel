package new

const MAIN_GO = `package main

import (
    "govel"
    "govel/app"
)

func main() {
    govel.Go(app.New())
}
`
