package router

type route interface{
	GetPath() string
	GetHandler() string
}