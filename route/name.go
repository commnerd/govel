package route

func (r *route) Name(name string) *route {
	r.name = name
	return r
}