package app

import (
	"github.com/commnerd/govel/config"
	_ "github.com/commnerd/govel/router"
)

func Bootstrap() {
	config.Inject(App)
	// router.Inject(app)
}
