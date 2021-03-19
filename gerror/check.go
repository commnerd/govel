package gerror

func Check(err interface{}) {
	if err != nil {
		Throw(err)
	}
}
