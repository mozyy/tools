package util

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	utils "github.com/mozyy/tools/utils/copy"
)

// CopyRimeFiles copy rime dir to user profile dir
func CopyRimeFiles() {
	var (
		rimeDir = os.Getenv("GOPATH") + "/src/github.com/mozyy/tools/rime/Rime/"
		distDir = os.Getenv("USERPROFILE") + "/AppData/Roaming/Rime/"
	)
	var wg sync.WaitGroup
	files, err := ioutil.ReadDir(rimeDir)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("copying ...")
	for _, file := range files {
		name := file.Name()
		wg.Add(1)
		go func() {
			_, err := utils.CopyFile(distDir+name, rimeDir+name)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("copying complete.")
}
