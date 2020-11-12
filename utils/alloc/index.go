package alloc

import (
	"bytes"
	"regexp"
	"strings"
)

func MakeAlloc() string {
	str := "1234567890"
	a := "12345"
	b := "aaaaa"
	return strings.Replace(str,a,b,1)
}
func MakePrevAlloc() string {
	str := "1234567890"
	a := "12345"
	b := "aaaaa"
	return string(bytes.Replace([]byte(str),[]byte(a),[]byte(b),1))
}

func RegTest(re *regexp.Regexp,str string){
	re.MatchString(str)
}