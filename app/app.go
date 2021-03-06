package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/commnerd/govel/config"
)

type application struct {
	children map[string]interface{}
}

var App = &application{
	children: make(map[string]interface{}),
}

func (a *application) Get(label string) interface{} {
	return a.children[label]
}

var Get = App.Get

func (a *application) Set(label string, item interface{}) {
	a.children[label] = item
}

var Set = App.Set

func (a *application) RootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

var RootPath = App.RootPath

func (a *application) GetConfig() *config.Config {
	switch config := a.Get("*config.Config").(type) {
	case *config.Config:
		return config
	default:
		panic("Something is wrong with your config type.")
	}
}

var GetConfig = App.GetConfig
