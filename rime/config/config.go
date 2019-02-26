package config

import (
	"strings"

	"github.com/mozyy/tools/rime/engin"
)

//signalDictName  = path + "/Rime/wubi091_signal.dict.yaml"
// specialDictName = path + "/Rime/wubi091_special.dict.yaml"
// wordDictName    = path + "/Rime/wubi091_word.dict.yaml"
// spreadDictName  = path + "/Rime/wubi091_spread.dict.yaml"

// Dicts array of dict
var Dicts = []engin.Dict{
	engin.Dict{ // special
		Result: []engin.CodeTableItem{},
		Name:   "/Rime/wubi091_special.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_special
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_special
version: "0.1"
sort: by_weight
...

`,
		Match: func(code, str string) bool {
			return strings.HasPrefix(str, "~")
		},
		BeforeAppend: func(code, str string, d engin.Dict) (string, string) {
			return code, string(str[1:])
		},
	},
	engin.Dict{ // spread
		Result: []engin.CodeTableItem{},
		Name:   "/Rime/wubi091_spread.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_spread
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_spread
version: "0.1"
sort: by_weight
...

`,
		Match: func(code, str string) bool {
			return strings.HasPrefix(code, "z")
		},
	},
	engin.Dict{ // signal
		Result: []engin.CodeTableItem{},
		Name:   "/Rime/wubi091_signal.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_signal
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

`,
		Match: func(code, str string) bool {
			return len([]rune(str)) < 2
		},
		BeforeAppend: func(code, str string, d engin.Dict) (string, string) {
			i, count, codeLen := len(d.Result), 0, len(code)
			if i > 2 && codeLen > 1 {
				lastCode := code[:codeLen-1]
				if d.Result[i-1].Code == lastCode && str != d.Result[i-1].Str {
					for i > 1 && d.Result[i-1].Code == lastCode {
						count++
						i--
					}
					if count < 3 {
						code = lastCode
					}
				}
			}
			return code, str
		},
	},
	engin.Dict{ // word
		Result: []engin.CodeTableItem{},
		Name:   "/Rime/wubi091_word.dict.yaml",
		ContentProfix: `# Rime dictionary: wubi091_word
# encoding: utf-8
#
# Original table author
# yangyue <yangyue@gmail.com>

---
name: wubi091_word
version: "0.1"
sort: by_weight
...

`,
		Match: func(code, str string) bool {
			return true
		},
	},
}
