package route

func (r *route) Only(methods []string) *route {
	r.only = methods
	return r
}