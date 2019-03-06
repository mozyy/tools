package config

import (
	"strings"

	"github.com/mozyy/tools/rime/engin"
)

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
...

`,
		Match: func(code, str string) bool {
			return strings.HasPrefix(str, "~")
		},
		BeforeAppend: func(code, str *string, d engin.Dict) {
			*str = string((*str)[1:])
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
...

`,
		Match: func(code, str string) bool {
			return len([]rune(str)) < 2
		},
		BeforeAppend: func(code, str *string, d engin.Dict) {
			transCode(code, str, d.Result)
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
...

`,
		Match: func(code, str string) bool {
			return true
		},
	},
}

// 阿	bs  -> 阿	bs
// 阱	bs  -> 阱	bs
// 耵	bsh -> 耵	bs
func transCode(code, str *string, c []engin.CodeTableItem) {

	for tCode := (*code)[:len(*code)-1]; len(tCode) > 0; tCode = tCode[:len(tCode)-1] {
		sameCode := findDict(c, func(c engin.CodeTableItem) bool {
			return c.Code == tCode
		})
		sameStr := findDict(c, func(c engin.CodeTableItem) bool {
			return c.Str == *str
		})

		if len(sameCode) < 3 {
			if len(sameStr) == 0 {
				// fmt.Printf("sameStr: %v, sameCode: %v, tCode: %s, str: %s \n", sameStr, sameCode, tCode, *str)
				*code = tCode
			} else if findMinCode(sameStr) > 1 || len(*code) != 4 {
				sameStrs := findDict(sameStr, func(c engin.CodeTableItem) bool {
					return c.Code == tCode
				})
				if len(sameStrs) == 0 {
					*code = tCode
				}
			} else {
				return
			}
		} else {
			return
		}
	}
}

func findDict(c []engin.CodeTableItem, f func(engin.CodeTableItem) bool) []engin.CodeTableItem {
	r := []engin.CodeTableItem{}
	for i := range c {
		if f(c[i]) {
			r = append(r, c[i])
		}
	}
	return r
}

func findMinCode(r []engin.CodeTableItem) int {
	i := len(r[0].Code)
	for _, v := range r {
		if len(v.Code) < i {
			i = len(v.Code)
		}
	}
	return i
}
