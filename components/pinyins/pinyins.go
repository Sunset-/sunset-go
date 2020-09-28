package pinyins

import (
	"bytes"
	"github.com/mozillazg/go-pinyin"
)

var PY_NORMAL = pinyin.Args{
	Style:     pinyin.NORMAL,
	Heteronym: false,
	Separator: "",
	Fallback:  nil,
}
var PY_LITE = pinyin.Args{
	Style:     pinyin.FIRST_LETTER,
	Heteronym: false,
	Separator: "",
	Fallback:  nil,
}

//全拼
func Pinyin(str string) string {
	var buff bytes.Buffer
	list := pinyin.Pinyin(str, PY_NORMAL)
	for _, v := range list {
		for _, v1 := range v {
			buff.WriteString(v1)
		}
	}
	return buff.String()
}

//简拼
func Acronym(str string) string {
	var buff bytes.Buffer
	list := pinyin.Pinyin(str, PY_LITE)
	for _, v := range list {
		for _, v1 := range v {
			buff.WriteString(v1)
		}
	}
	return buff.String()
}

//全拼+简拼
func PinyinAcronym(str string) string {
	return Pinyin(str) + "_" + Acronym(str)
}
