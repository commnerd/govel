package gerror

type gError interface{
	Error() error
}