package engin

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

// Run is begin engin to generate rime dict
func Run(path string, dicts []Dict) error {
	log.Println("converting ...")
	defer log.Println("converting complete.")

	nativeCodeTable := path + "/点儿词库1901版/码表.txt"
	file, err := os.Open(nativeCodeTable)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	codeTableRg := regexp.MustCompile(`^(\S+)\s(.*)$`)

	// limit := 20
	for scanner.Scan() {
		b := scanner.Text()
		matchs := codeTableRg.FindAllStringSubmatch(b, -1)
		for _, match := range matchs {
			strs := strings.Split(match[2], " ")
			for _, str := range strs {
				code := match[1]

				for i, dict := range dicts {
					if dict.Match(code, str) {
						if dict.BeforeAppend != nil {
							code, str = dict.BeforeAppend(code, str, dict)
						}
						dicts[i].Append(code, str)
						break
					}
				}
			}
		}
		// limit--
		// if limit < 0 {
		// 	break
		// }
	}
	wg := sync.WaitGroup{}
	for _, dict := range dicts {
		wg.Add(1)
		go func(dict Dict) {
			distName := path + dict.Name
			content := []byte(dict.String())
			err := ioutil.WriteFile(distName, content, 32)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(dict)
	}
	wg.Wait()

	return nil
}
