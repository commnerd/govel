package server

import (
	"net/http"
	"strconv"

	"github.com/commnerd/govel/config"
	"github.com/commnerd/govel/router"
)

func Serve() {
	config := config.Get()

	router := router.Get()

	hostAddr := ":" + strconv.Itoa(config.GetInt("main.port"))

	http.ListenAndServe(hostAddr, router)

}
