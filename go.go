package govel

import (
	"fmt"
	"log"
	"os"
)

func Go() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}
