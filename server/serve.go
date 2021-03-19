package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/commnerd/govel/config"
	"github.com/commnerd/govel/gerror"
	"github.com/commnerd/govel/router"
)

func Serve() {
	router := app.Get("*router.Router").(*router.Router)

	hostAddr := ":" + strconv.Itoa(config.Get().GetInt("main.port"))

	_, err := fmt.Printf("Serving \"%s\"", hostAddr)
	gerror.Check(err)

	http.ListenAndServe(hostAddr, router)

}
