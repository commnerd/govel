package route

func Resource(path, handler string) *routeStruct {
	return craft(path, handler, "Resource")
}
