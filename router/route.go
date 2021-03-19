package router

type route interface {
	GetPath() string
	GetController() string
	GetHandler(string) string
}
