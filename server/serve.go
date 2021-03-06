package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/commnerd/govel/app"
)

func Serve() {
	config := app.GetConfig()

	hostAddr := ":" + strconv.Itoa(config.GetInt("main.port"))

	// router := app.Get("router")

	// http.ListenAndServe(":8090", router)

	fmt.Println(fmt.Sprintf("Serving: '%s'", hostAddr))

	http.ListenAndServe(hostAddr, nil)

}
