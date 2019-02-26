package util

import (
	"bufio"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
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

type codeTables struct {
	singleCodeTables  []codeTableItem
	wordCodeTables    []codeTableItem
	specialCodeTables []codeTableItem
	spreadCodeTables  []codeTableItem
}

//GenerateRime will generate rime dict file
func GenerateRime() {
	var (
		path            = build.Default.GOPATH + "/src/github.com/mozyy/tools/rime"
		name            = path + "/点儿词库1901版/码表.txt"
		signalDictName  = path + "/Rime/wubi091_signal.dict.yaml"
		specialDictName = path + "/Rime/wubi091_special.dict.yaml"
		wordDictName    = path + "/Rime/wubi091_word.dict.yaml"
		spreadDictName  = path + "/Rime/wubi091_spread.dict.yaml"
		line            int
	)
	signalContent :=
		`# Rime dictionary: wubi091_signal
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_signal
version: "0.1"
sort: by_weight
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
	specialContent :=
		`# Rime dictionary: wubi091_special
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_special
version: "0.1"
sort: by_weight
...

`
	wordContent :=
		`# Rime dictionary: wubi091_word
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_word
version: "0.1"
sort: by_weight
...

`
	spreadContent :=
		`# Rime dictionary: wubi091_spread
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_spread
version: "0.1"
sort: by_weight
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
			codeTables := codeTables{
				[]codeTableItem{},
				[]codeTableItem{},
				[]codeTableItem{},
				[]codeTableItem{},
			}
			for i, str := range strs {
				code := match[1]

				// 过滤zv开头生僻字
				if !strings.HasPrefix(code, "zv") {

					if strings.HasPrefix(str, "~") { // 特殊字
						codeTables.specialCodeTables = append(codeTables.specialCodeTables,
							codeTableItem{str[1:], code, weight - i})
					} else if strings.HasPrefix(code, "z") { // z 开头
						codeTables.spreadCodeTables = append(codeTables.spreadCodeTables,
							codeTableItem{str, code, weight - i})
					} else if len([]rune(str)) < 2 { // 单字
						codeTables.singleCodeTables = append(codeTables.singleCodeTables,
							codeTableItem{str, code, weight - i})
					} else { // 词语
						codeTables.wordCodeTables = append(codeTables.wordCodeTables,
							codeTableItem{str, code, weight - i})
					}
				}
			}

			sort.Sort(codeTable(codeTables.specialCodeTables))
			sort.Sort(codeTable(codeTables.singleCodeTables))
			sort.Sort(codeTable(codeTables.wordCodeTables))
			sort.Sort(codeTable(codeTables.spreadCodeTables))
			for _, item := range codeTables.specialCodeTables {
				specialContent += item.String()
			}
			for _, item := range codeTables.singleCodeTables {
				signalContent += item.String()
			}
			for _, item := range codeTables.wordCodeTables {
				wordContent += item.String()
			}
			for _, item := range codeTables.spreadCodeTables {
				spreadContent += item.String()
			}
		}
		// limit--
		// if limit < 0 {
		// 	break
		// }
	}
	// fmt.Println(text)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		writeFile(specialDictName, specialContent)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		writeFile(signalDictName, signalContent)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		writeFile(wordDictName, wordContent)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		writeFile(spreadDictName, spreadContent)
		wg.Done()
	}()
	wg.Wait()
	log.Println("converting complete.")
}

func writeFile(name, content string) {
	err := ioutil.WriteFile(name, []byte(content), 32)
	if err != nil {
		log.Println(err)
	}
}
