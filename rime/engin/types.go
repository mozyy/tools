package engin

import "fmt"

// Dict is engin rime dict config
type Dict struct {
	Result              []CodeTableItem
	Name, ContentProfix string
	Match               func(code, str string) bool
	BeforeAppend        func(code, str string, d Dict) (string, string)
}

// Append is append CodeTableItem to Dict
func (d *Dict) Append(code, str string) {
	d.Result = append(d.Result, CodeTableItem{str, code})
}

// String is func
func (d *Dict) String() string {
	str := d.ContentProfix
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
