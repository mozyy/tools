package engin

import (
	"bufio"
	"go/build"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestRun(t *testing.T) {
	var (
		path    = build.Default.GOPATH + "/src/github.com/mozyy/tools/rime/engin/test/"
		rimeDir = build.Default.GOPATH + "/src/github.com/mozyy/tools/rime/engin/test/Rime/"
	)
	err := Run(path, dicts)
	if err != nil {
		t.Errorf("[engin run error]:%s\n", err.Error())
	}

	files, err := ioutil.ReadDir(rimeDir)
	if err != nil {
		t.Errorf("[engin ReadDir error]:%s\n", err.Error())
	}
	wg := sync.WaitGroup{}
	for _, fileInfo := range files {
		name := fileInfo.Name()
		wg.Add(1)
		go func() {
			generatePaht := path + "temp/" + name
			defer func() {
				err = os.Remove(generatePaht)
				if err != nil {
					t.Errorf("[engin Remove generate]:%s\n", err.Error())
				}
			}()
			generate, err := os.Open(generatePaht)
			if err != nil {
				t.Errorf("[engin Open generate]:%s\n", err.Error())
			}
			defer generate.Close()
			origin, err := os.Open(rimeDir + name)
			if err != nil {
				t.Errorf("[engin Open origin]:%s\n", err.Error())
			}
			defer origin.Close()
			geneScan := bufio.NewScanner(generate)
			origScan := bufio.NewScanner(origin)
			for origScan.Scan() {
				if !geneScan.Scan() {
					t.Errorf("[generate file error]: file: %s", generatePaht)
					break
				}
				geneText, origText := geneScan.Text(), origScan.Text()
				if geneText != origText {
					t.Errorf("[generate file error]: file: %s, want: %s, has: %s", generatePaht, origText, geneText)
					break
				}
			}
			if geneScan.Scan() {
				t.Errorf("[origin file error]: file: %s", generatePaht)
			}

			wg.Done()
		}()
	}
	wg.Wait()
}

var dicts = []Dict{
	Dict{ // special
		Result: []CodeTableItem{},
		Name:   "/temp/wubi091_special.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_special
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_special
version: "0.1"
...

`,
		Match: func(code, str string) bool {
			return strings.HasPrefix(str, "~")
		},
		BeforeAppend: func(code, str *string, d Dict) {
			*str = string((*str)[1:])
		},
	},
	Dict{ // spread
		Result: []CodeTableItem{},
		Name:   "/temp/wubi091_spread.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_spread
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_spread
version: "0.1"
...

`,
		Match: func(code, str string) bool {
			return strings.HasPrefix(code, "z")
		},
	},
	Dict{ // signal
		Result: []CodeTableItem{},
		Name:   "/temp/wubi091_signal.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_signal
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_signal
version: "0.1"
...

`,
		Match: func(code, str string) bool {
			return len([]rune(str)) < 2
		},
		BeforeAppend: func(code, str *string, d Dict) {
			i, count, codeLen := len(d.Result), 0, len(*code)
			// 阿	bs  -> 阿	bs
			// 阱	bs  -> 阱	bs
			// 耵	bsh -> 耵	bs
			if i > 2 && codeLen > 1 {
				lastCode := (*code)[:codeLen-1]
				if d.Result[i-1].Code == lastCode && *str != d.Result[i-1].Str {
					for i > 0 && d.Result[i-1].Code == lastCode {
						count++
						i--
					}
					if count < 3 {
						*code = lastCode
					}
				}
			}
		},
	},
	Dict{ // word
		Result: []CodeTableItem{},
		Name:   "/temp/wubi091_word.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_word
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_word
version: "0.1"
...

`,
		Match: func(code, str string) bool {
			return true
		},
	},
}
