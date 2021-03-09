package app

func GetConfig() interface {
	Get(string) interface{}
	Set(string) interface{}
	GetInt(string) int
} {
	return Get("config").(interface {
		Get(string) interface{}
		Set(string) interface{}
		GetInt(string) int
	})
}
