package util

import (
	"io/ioutil"
	"log"
	"sync"

	utils "github.com/mozyy/tools/utils/copy"
)

// CopyRimeFiles copy rime dir to user profile dir
func CopyRimeFiles() {
	var (
		rimeDir = ".\\Rime\\"
		distDir = "C:\\Users\\PC028\\AppData\\Roaming\\Rime\\"
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
			b, err := utils.CopyFile(distDir+name, rimeDir+name)
			if err != nil {
				log.Println(err)
			}
			log.Printf("src: %s, name: %s, length: %d", rimeDir+name, distDir+name, b)
			wg.Done()
		}()
	}
	wg.Wait()
}
