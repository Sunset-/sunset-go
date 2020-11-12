package casts

import (
	"github.com/Sunset-/sunset-go/utils/alloc"
	"regexp"
	"testing"
)

func BenchmarkTestAbc(b *testing.B){
	re,_ := regexp.Compile("^[陕禁微风感觉一天就正常也弄下]\\d{8}$")

	for i:=0;i<b.N;i++{
		alloc.RegTest(re,"陕A123123")
	}
}