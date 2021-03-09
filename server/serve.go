package server

import (
	"net/http"
	"strconv"

	"github.com/commnerd/govel/app"
	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/router"
)

func Serve() {
	config := app.GetConfig()

	hostAddr := ":" + strconv.Itoa(config.GetInt("main.port"))

	router, ok := app.Get("router").(*router.Router)
	if !ok {
		gerror.Throw("Router not properly implemented.")
	}

	http.ListenAndServe(hostAddr, router)

}
