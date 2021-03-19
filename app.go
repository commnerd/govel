package govel

type app interface {
	Register(i interface{}) error
	Path(string) string
}
