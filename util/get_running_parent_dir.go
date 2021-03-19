package util

import (
	"os"
	"path/filepath"

	"github.com/commnerd/govel/gerror"
)

func GetRunningParentDir() string {
	runDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	gerror.Check(err)

	cwd, err := os.Getwd()
	gerror.Check(err)

	if string(runDir[0:4]) == "/tmp" && cwd != runDir {
		return cwd
	}

	return runDir
}
