package app

import "os"

func (a *application) Path(path string) string {
	if string(path[0]) != string(os.PathSeparator) {
		path = string(os.PathSeparator) + path
	}
	return BasePath() + path
}
