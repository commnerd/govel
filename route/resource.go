package route

func Resource(path, handler string) *route {
	return craft(path, handler, "Resource")
}
