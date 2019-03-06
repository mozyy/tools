package engin

import (
	"fmt"
	"sort"
)

type codeTableSlice []CodeTableItem

func (c codeTableSlice) Len() int {
	return len(c)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (c codeTableSlice) Less(i, j int) bool {
	return c[i].Code < c[j].Code
}

// Swap swaps the elements with indexes i and j.
func (c codeTableSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Dict is engin rime dict config
type Dict struct {
	Result              []CodeTableItem
	Name, ContentProfix string
	Match               func(code, str string) bool
	BeforeAppend        func(code, str *string, d Dict)
}

// Append is append CodeTableItem to Dict
func (d *Dict) Append(code, str string) {
	d.Result = append(d.Result, CodeTableItem{str, code})
}

// String is func
func (d *Dict) String() string {
	str := d.ContentProfix
	sort.Stable(codeTableSlice(d.Result))
	for _, v := range d.Result {
		str += v.String()
	}
	return str
}

// CodeTableItem is rime code table item
type CodeTableItem struct {
	Str  string
	Code string
}

func (i CodeTableItem) String() string {
	return fmt.Sprintf("%s\t%s\n", i.Str, i.Code)
}
