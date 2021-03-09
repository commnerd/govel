package router

func (r *Router) Register(i route) error {
	r.Set(i.GetPath(), i)
	return nil
}

var Register = instance.Register
