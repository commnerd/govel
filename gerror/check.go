package gerror

func Check(err error) {
	if err != nil {
		Throw(err)
	}
}