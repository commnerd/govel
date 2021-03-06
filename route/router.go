package route

import (
	"fmt"
	"log"
	"os"
)

type Router struct{}

func (r *Router) Bootstrap() {
	printRoutes()
}

func printRoutes() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(path)
}
