package app

func (a *application) Get(label string) interface{} {
	return a.Viper.Get(label)
}
