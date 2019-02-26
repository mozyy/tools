package main

import (
	"go/build"

	"github.com/mozyy/tools/rime/config"
	"github.com/mozyy/tools/rime/engin"
	"github.com/mozyy/tools/rime/util"
)

func main() {
	var (
		path = build.Default.GOPATH + "/src/github.com/mozyy/tools/rime"
	)
	engin.Run(path, config.Dicts)
	// util.GenerateRime()
	util.CopyRimeFiles()

}
