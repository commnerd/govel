package config

type iApp interface {
	Set(string, interface{})
	RootPath() string
}

var app iApp
