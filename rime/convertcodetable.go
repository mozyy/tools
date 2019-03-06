package main

import (
	"flag"
	"go/build"
	"log"

	"github.com/mozyy/tools/rime/config"
	"github.com/mozyy/tools/rime/engin"
	"github.com/mozyy/tools/rime/util"
)

func main() {
	var (
		path              = build.Default.GOPATH + "/src/github.com/mozyy/tools/rime"
		help, parse, copy bool
	)
	flag.BoolVar(&help, "h", false, "this help")
	flag.BoolVar(&parse, "np", false, "not parse code table")
	flag.BoolVar(&copy, "nc", false, "not copy to userfile dir")
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if !parse {
		err := engin.Run(path, config.Dicts)
		if err != nil {
			log.Panicln(err)
		}
	}
	if !copy {
		// util.GenerateRime()
		util.CopyRimeFiles()
	}
}
