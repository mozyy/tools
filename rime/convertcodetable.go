package main

import (
	"go/build"
	"log"

	"github.com/mozyy/tools/rime/config"
	"github.com/mozyy/tools/rime/engin"
	"github.com/mozyy/tools/rime/util"
)

func main() {
	var (
		path = build.Default.GOPATH + "/src/github.com/mozyy/tools/rime"
	)
	err := engin.Run(path, config.Dicts)
	if err != nil {
		log.Panicln(err)
	}
	// util.GenerateRime()
	util.CopyRimeFiles()

}
