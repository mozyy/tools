package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/mozyy/tools/rime/util"
)

type codeTableItem struct {
	str    string
	code   string
	weight int
}

func (i codeTableItem) String() string {
	return fmt.Sprintf("%s\t%s\t%d\n", i.str, i.code, i.weight)
}

type codeTable []codeTableItem

func (p codeTable) Len() int           { return len(p) }
func (p codeTable) Less(i, j int) bool { return i < j }
func (p codeTable) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	var (
		name    = ".\\点儿词库1901版\\码表.txt"
		txtName = ".\\Rime\\wubi091.dict.yaml"
		text    string
		line    int
	)
	text =
		`# Rime dictionary: wubi091
# encoding: utf-8
#
# Changelog
#
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091
version: "0.1"
sort: by_weight
columns:
  - text
  - code
  - weight
  # - stem
encoder:
  exclude_patterns:
    - '^z.*$'
  rules:
    - length_equal: 2
      formula: "AaAbBaBb"
    - length_equal: 3
      formula: "AaBaCaCb"
    - length_in_range: [4, 10]
      formula: "AaBaCaZa"
...

`
	log.Println("converting ...")
	file, err := os.Open(name)
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)
	codeTableRg := regexp.MustCompile(`^(\S+)\s(.*)$`)

	// limit := 20
	for scanner.Scan() {
		b := scanner.Text()
		line++
		matchs := codeTableRg.FindAllStringSubmatch(b, -1)
		for _, match := range matchs {
			strs := strings.Split(match[2], " ")
			weight := line * 10
			codeTables := []codeTableItem{}
			for i, str := range strs {
				code := match[1]

				if !strings.HasPrefix(code, "zv") {
					if strings.HasPrefix(str, "~") {
						*&str = "# " + str
					}
					item := codeTableItem{str, code, weight - i}
					codeTables = append(codeTables, item)
				}
			}
			sort.Sort(codeTable(codeTables))
			for _, item := range codeTables {
				text += item.String()
			}
		}
		// limit--
		// if limit < 0 {
		// 	break
		// }
	}
	// fmt.Println(text)
	err = ioutil.WriteFile(txtName, []byte(text), 32)
	if err != nil {
		log.Println(err)
	}
	util.CopyRimeFiles()
}
