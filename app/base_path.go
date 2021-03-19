package app

import "github.com/commnerd/govel/util"

func BasePath() string {
	return util.GetRunningParentDir()
}
