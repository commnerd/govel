package route

func (r *routeStruct) Only(methods []string) *routeStruct {
	r.only = methods
	return r
}
